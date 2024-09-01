<template>
  <div>
    <div class="gva-search-box">
      <el-form
        ref="elSearchFormRef"
        :inline="true"
        :model="searchInfo"
        :rules="searchRule"
        class="demo-form-inline"
        @keyup.enter="onSubmit"
      >
        <el-form-item label="创建日期" prop="createdAt">
          <template #label>
            <span>
              创建日期
              <el-tooltip
                content="搜索范围是开始日期（包含）至结束日期（不包含）"
              >
                <el-icon><QuestionFilled /></el-icon>
              </el-tooltip>
            </span>
          </template>
          <el-date-picker
            v-model="searchInfo.startCreatedAt"
            :disabled-date="
              (time) =>
                searchInfo.endCreatedAt
                  ? time.getTime() > searchInfo.endCreatedAt.getTime()
                  : false
            "
            placeholder="开始日期"
            type="datetime"
          />
          —
          <el-date-picker
            v-model="searchInfo.endCreatedAt"
            :disabled-date="
              (time) =>
                searchInfo.startCreatedAt
                  ? time.getTime() < searchInfo.startCreatedAt.getTime()
                  : false
            "
            placeholder="结束日期"
            type="datetime"
          />
        </el-form-item>

        <el-form-item>
          <el-button icon="search" type="primary" @click="onSubmit">
            查询
          </el-button>
          <el-button icon="refresh" @click="onReset"> 重置 </el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button icon="plus" type="primary" @click="openDialog">
          新增1
        </el-button>
      </div>
      <el-table
        ref="multipleTable"
        :data="tableData"
        row-key="ID"
        style="width: 100%"
        tooltip-effect="dark"
        @selection-change="handleSelectionChange"
      >
        <el-table-column align="left" label="日期" width="180">
          <template #default="scope">
            {{ formatDate(scope.row.CreatedAt) }}
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="交易人nickName"
          prop="user.nickName"
          width="120"
        />
        <el-table-column
          align="left"
          label="交易人userName"
          prop="user.userName"
          width="120"
        />
        <el-table-column
          align="left"
          label="交易人手机号"
          prop="user.phone"
          width="120"
        />
        <el-table-column
          align="left"
          label="交易前金额"
          prop="preTransactionAmount"
          width="120"
        />
        <el-table-column
          align="left"
          label="交易金额"
          prop="transactionAmount"
          width="120"
        />
        <el-table-column
          align="left"
          label="交易后金额"
          prop="postTransactionAmount"
          width="120"
        />
        <el-table-column
          align="left"
          label="微信订单号"
          prop="wechatOrderId"
          show-overflow-tooltip
          width="320"
        />

        <el-table-column
          align="left"
          label="货单号"
          prop="billNumber"
          show-overflow-tooltip
          width="320"
        />
        <el-table-column
          align="left"
          label="备注"
          prop="remark"
          show-overflow-tooltip
          width="320"
        />
        &lt;!&ndash;
        <el-table-column
          align="left"
          fixed="right"
          label="操作"
          min-width="240"
        >
          <template #default="scope">
            <el-button
              class="table-button"
              icon="edit"
              link
              type="primary"
              @click="updateTransactionDetailsFunc(scope.row)"
            >
              变更
            </el-button>
            <el-button
              icon="delete"
              link
              type="primary"
              @click="deleteRow(scope.row)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
        &ndash;&gt;
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-drawer
      v-model="dialogFormVisible"
      :before-close="closeDialog"
      :show-close="false"
      destroy-on-close
      size="800"
    >
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{ type === "create" ? "添加" : "修改" }}</span>
          <div>
            <el-button type="primary" @click="enterDialog"> 确 定 </el-button>
            <el-button @click="closeDialog"> 取 消 </el-button>
          </div>
        </div>
      </template>

      <el-form
        ref="elFormRef"
        :model="formData"
        :rules="rule"
        label-position="top"
        label-width="80px"
      >
        <el-form-item label="用户" prop="user">
          <el-select
            v-model="formData.user"
            :clearable="true"
            filterable
            placeholder="请选择用户"
            value-key="ID"
          >
            <el-option
              v-for="(item, index) in userOption"
              :key="index"
              :label="item.userName"
              :value="item"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="交易金额:" prop="transactionAmount">
          <el-input-number
            v-model="formData.transactionAmount"
            :clearable="true"
            :precision="2"
            style="width: 100%"
          />
        </el-form-item>
      </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createTransactionDetails,
  deleteTransactionDetails,
  deleteTransactionDetailsByIds,
  findTransactionDetails,
  getTransactionDetailsList,
  updateTransactionDetails,
} from "@/api/transaction/transactionDetails";

// 全量引入格式化工具 请按需保留
import { formatDate } from "@/utils/format";
import { ElMessage, ElMessageBox } from "element-plus";
import { reactive, ref } from "vue";
import { getUserList } from "@/api/user";

defineOptions({
  name: "TransactionDetails",
});

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  preTransactionAmount: 0,
  transactionAmount: 0,
  postTransactionAmount: 0,
  user: {},
});

// 验证规则
const rule = reactive({});

const searchRule = reactive({
  createdAt: [
    {
      validator: (rule, value, callback) => {
        if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
          callback(new Error("请填写结束日期"));
        } else if (
          !searchInfo.value.startCreatedAt &&
          searchInfo.value.endCreatedAt
        ) {
          callback(new Error("请填写开始日期"));
        } else if (
          searchInfo.value.startCreatedAt &&
          searchInfo.value.endCreatedAt &&
          (searchInfo.value.startCreatedAt.getTime() ===
            searchInfo.value.endCreatedAt.getTime() ||
            searchInfo.value.startCreatedAt.getTime() >
              searchInfo.value.endCreatedAt.getTime())
        ) {
          callback(new Error("开始日期应当早于结束日期"));
        } else {
          callback();
        }
      },
      trigger: "change",
    },
  ],
});

const elFormRef = ref();
const elSearchFormRef = ref();

// =========== 表格控制部分 ===========
const page = ref(1);
const total = ref(0);
const pageSize = ref(10);
const tableData = ref([]);
const searchInfo = ref({});

// 重置
const onReset = () => {
  searchInfo.value = {};
  getTableData();
};

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async (valid) => {
    if (!valid) return;
    page.value = 1;
    pageSize.value = 10;
    getTableData();
  });
};

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val;
  getTableData();
};

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val;
  getTableData();
};

// 查询
const getTableData = async () => {
  const table = await getTransactionDetailsList({
    page: page.value,
    pageSize: pageSize.value,
    ...searchInfo.value,
  });
  if (table.code === 0) {
    tableData.value = table.data.list;
    total.value = table.data.total;
    page.value = table.data.page;
    pageSize.value = table.data.pageSize;
  }
};

getTableData();

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () => {};

// 获取需要的字典 可能为空 按需保留
setOptions();

// 多选数据
const multipleSelection = ref([]);
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val;
};

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm("确定要删除吗?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  }).then(() => {
    deleteTransactionDetailsFunc(row);
  });
};

// 多选删除
const onDelete = async () => {
  ElMessageBox.confirm("确定要删除吗?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  }).then(async () => {
    const IDs = [];
    if (multipleSelection.value.length === 0) {
      ElMessage({
        type: "warning",
        message: "请选择要删除的数据",
      });
      return;
    }
    multipleSelection.value &&
      multipleSelection.value.map((item) => {
        IDs.push(item.ID);
      });
    const res = await deleteTransactionDetailsByIds({ IDs });
    if (res.code === 0) {
      ElMessage({
        type: "success",
        message: "删除成功",
      });
      if (tableData.value.length === IDs.length && page.value > 1) {
        page.value--;
      }
      getTableData();
    }
  });
};

// 行为控制标记（弹窗内部需要增还是改）
const type = ref("");

// 更新行
const updateTransactionDetailsFunc = async (row) => {
  const res = await findTransactionDetails({ ID: row.ID });
  type.value = "update";
  if (res.code === 0) {
    formData.value = res.data.retd;
    dialogFormVisible.value = true;
  }
};

// 删除行
const deleteTransactionDetailsFunc = async (row) => {
  const res = await deleteTransactionDetails({ ID: row.ID });
  if (res.code === 0) {
    ElMessage({
      type: "success",
      message: "删除成功",
    });
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--;
    }
    getTableData();
  }
};

// 弹窗控制标记
const dialogFormVisible = ref(false);

// 打开弹窗
const openDialog = () => {
  type.value = "create";
  dialogFormVisible.value = true;
};

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false;
  formData.value = {
    preTransactionAmount: 0,
    transactionAmount: 0,
    postTransactionAmount: 0,
  };
};
// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return;
    let res;
    switch (type.value) {
      case "create":
        res = await createTransactionDetails({
          ...formData.value,
          userId: formData.value.user.userId,
        });
        break;
      case "update":
        res = await updateTransactionDetails(formData.value);
        break;
      default:
        res = await createTransactionDetails(formData.value);
        break;
    }
    if (res.code === 0) {
      ElMessage({
        type: "success",
        message: "创建/更改成功",
      });
      closeDialog();
      getTableData();
    }
  });
};

const userOption = ref([]);
const queryUserList = async () => {
  const res = await getUserList({
    page: 1,
    pageSize: 1000,
  });
  userOption.value = res.data.list;
};
queryUserList();
</script>

<style></style>
