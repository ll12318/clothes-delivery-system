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
          ></el-date-picker>
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
          ></el-date-picker>
        </el-form-item>

        <el-form-item label="路线名称" prop="routeName">
          <el-input v-model="searchInfo.routeName" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="司机" prop="routeName">
          <el-select
            v-model="searchInfo.userIds"
            clearable
            collapse-tags
            filterable
            multiple
            placeholder="请选择司机"
          >
            <el-option
              v-for="(item, index) in userOption"
              :key="index"
              :label="item.nickName"
              :value="item.ID"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="加急" prop="urgent">
          <el-select v-model="searchInfo.urgent" clearable filterable>
            <el-option
              v-for="item in [
                { label: '是', value: true },
                { label: '否', value: false },
              ]"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button icon="search" type="primary" @click="onSubmit"
            >查询</el-button
          >
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button icon="plus" type="primary" @click="openDialog"
          >新增</el-button
        >
        <el-button
          :disabled="!multipleSelection.length"
          icon="delete"
          style="margin-left: 10px"
          @click="onDelete"
          >删除</el-button
        >
      </div>
      <el-table
        ref="multipleTable"
        :data="tableData"
        row-key="ID"
        style="width: 100%"
        tooltip-effect="dark"
        @selection-change="handleSelectionChange"
        @sort-change="sortChange"
      >
        <el-table-column type="selection" width="55" />

        <el-table-column align="left" label="日期" width="180">
          <template #default="scope">{{
            formatDate(scope.row.CreatedAt)
          }}</template>
        </el-table-column>

        <el-table-column
          align="left"
          label="路线名称"
          prop="routeName"
          sortable
          width="120"
        />
        <el-table-column
          align="left"
          label="档口"
          show-overflow-tooltip
          width="auto"
        >
          <template #default="scope">{{
            scope.row.stalls.map((item) => item.stall).join(",")
          }}</template>
        </el-table-column>

        <el-table-column
          align="left"
          label="司机"
          prop="user.nickName"
          width="120"
        />
        <el-table-column align="left" label="加急" prop="urgent" width="120">
          <template  #default="scope">
            <div>
              {{ scope.row.urgent? '加急':'非加急' }}
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="备注" prop="remarks" width="120" />
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
              @click="updateRouteFunc(scope.row)"
              >变更</el-button
            >
            <el-button
              icon="delete"
              link
              type="primary"
              @click="deleteRow(scope.row)"
              >删除</el-button
            >
          </template>
        </el-table-column>
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
        <div class="flex items-center justify-between">
          <span class="text-lg">{{ type === "create" ? "添加" : "修改" }}</span>
          <div>
            <el-button type="primary" @click="enterDialog">确 定</el-button>
            <el-button @click="closeDialog">取 消</el-button>
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
        <el-form-item label="路线名称:" prop="routeName">
          <el-input
            v-model="formData.routeName"
            :clearable="true"
            placeholder="请输入路线名称"
          />
        </el-form-item>
        <el-form-item label="加急:" prop="urgent">
          <el-switch v-model="formData.urgent" @change="urgentChange" />
        </el-form-item>
        <el-form-item
          :label="formData.urgent ? '加急档口' : '档口'"
          prop="stalls"
        >
          <el-select
            v-model="formData.stalls"
            :clearable="true"
            :placeholder="formData.urgent ? '请选择加急档口' : '请选择档口'"
            filterable
            multiple
            value-key="ID"
          >
            <el-option
              v-for="(item, index) in stallOption"
              :key="index"
              :label="item.stall"
              :value="item"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="司机" prop="user">
          <el-select
            v-model="formData.user"
            :clearable="true"
            filterable
            placeholder="请选择司机"
            value-key="ID"
          >
            <el-option
              v-for="(item, index) in userOption"
              :key="index"
              :label="item.nickName"
              :value="item"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="备注:" prop="remarks">
          <el-input
            v-model="formData.remarks"
            :clearable="true"
            placeholder="请输入备注"
          />
        </el-form-item>
      </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createRoute,
  deleteRoute,
  deleteRouteByIds,
  findRoute,
  getRouteList,
  updateRoute,
} from "@/api/dataConfig/route";
import { getStallList } from "@/api/dataConfig/stall";
import { getUserList } from "@/api/user";
// 全量引入格式化工具 请按需保留
import { formatDate } from "@/utils/format";
import { ElMessage, ElMessageBox } from "element-plus";
import { reactive, ref } from "vue";

defineOptions({
  name: "Route",
});

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  routeName: "",
  remarks: "",
  stalls: null,
  urgent: false,
});

const stallOption = ref([]);

const queryStallList = async () => {
  const res = await getStallList({
    page: 1,
    pageSize: 1000,
    urgent: formData.value.urgent,
    filterOccupancy: true,
  });
  stallOption.value = res.data.list;
};

const userOption = ref([]);

const queryUserList = async () => {
  const res = await getUserList({
    page: 1,
    pageSize: 1000,
    authorityIds: [999],
  });
  userOption.value = res.data.list;
};
queryUserList();
// 验证规则
const rule = reactive({
  routeName: [
    {
      required: true,
      message: "",
      trigger: ["input", "blur"],
    },
    {
      whitespace: true,
      message: "不能只输入空格",
      trigger: ["input", "blur"],
    },
  ],
});

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
// 排序
const sortChange = ({ prop, order }) => {
  const sortMap = {
    routeName: "route_name",
  };

  let sort = sortMap[prop];
  if (!sort) {
    sort = prop.replace(/[A-Z]/g, (match) => `_${match.toLowerCase()}`);
  }

  searchInfo.value.sort = sort;
  searchInfo.value.order = order;
  getTableData();
};

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
  const table = await getRouteList({
    page: page.value,
    pageSize: pageSize.value,
    ...searchInfo.value,
    userIds: JSON.stringify(searchInfo.value.userIds),
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
    deleteRouteFunc(row);
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
    const res = await deleteRouteByIds({ IDs });
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
const updateRouteFunc = async (row) => {
  const res = await findRoute({ ID: row.ID });
  type.value = "update";
  if (res.code === 0) {
    formData.value = res.data.reroute;
    dialogFormVisible.value = true;
    await queryStallList();
    await queryUserList();
    stallOption.value.push(...row.stalls);
  }
};

// 删除行
const deleteRouteFunc = async (row) => {
  const res = await deleteRoute({ ID: row.ID });
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
  queryStallList();
  queryUserList();
};

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false;
  formData.value = {
    routeName: "",
    remarks: "",
    stalls: null,
  };
  stallOption.value = [];
  userOption.value = [];
};
// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return;
    let res;
    switch (type.value) {
      case "create":
        if (formData.value.urgent== undefined) {
          formData.value.urgent = false
        }
        res = await createRoute(formData.value);
        break;
      case "update":
        res = await updateRoute({
          ...formData.value,
          userId: formData.value.user.ID,
        });
        break;
      default:
        res = await createRoute(formData.value);
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

const urgentChange = () => queryStallList();
</script>

<style></style>
