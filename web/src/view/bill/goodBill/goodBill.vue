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

        <el-form-item label="单据编号" prop="billNumber">
          <el-input
            v-model="searchInfo.billNumber"
            placeholder="请输入单据编号"
          />
        </el-form-item>

        <el-form-item
          v-if="btnAuth.takeGoodPeopleInp"
          label="下单人"
          prop="createdBy"
        >
          <el-select
            v-model="searchInfo.createdBy"
            clearable
            filterable
            placeholder="请选择下单人"
          >
            <el-option
              v-for="(item, index) in createdOption"
              :key="index"
              :label="item.nickName"
              :value="item.ID"
            />
          </el-select>
        </el-form-item>
        <el-form-item
          v-if="btnAuth.takeGoodPeopleInp"
          label="拿货人"
          prop="takeGoodPeopleId"
        >
          <el-select
            v-model="searchInfo.takeGoodPeopleId"
            clearable
            filterable
            placeholder="请选择拿货人"
          >
            <el-option
              v-for="(item, index) in userOption"
              :key="index"
              :label="item.nickName"
              :value="item.ID"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="拿货状态" prop="finishStatus">
          <el-select v-model="searchInfo.finishStatus" clearable filterable>
            <el-option
              v-for="item in [
                { label: '是', value: true },
                { label: '否', value: false },
              ]"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="是否到达集散点" prop="driverVerify">
          <el-select v-model="searchInfo.driverVerify" clearable filterable>
            <el-option
              v-for="item in [
                { label: '是', value: true },
                { label: '否', value: false },
              ]"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="下单设备" prop="device">
          <el-select v-model="searchInfo.device" clearable filterable>
            <el-option
              v-for="item in [
                { label: '小程序', value: '0' },
                { label: '网页端', value: '1' },
              ]"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="审核订单" prop="isManual">
          <el-select v-model="searchInfo.isManual" clearable filterable>
            <el-option
              v-for="item in [
                { label: '非审核单', value: '0' },
                { label: '审核单', value: '1' },
              ]"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
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
      <div
        class="gva-btn-list"
        style="display: flex; justify-content: space-between"
      >
        <div>
          <el-button icon="plus" type="primary" @click="openDialog">
            新增
          </el-button>
          <el-button
            v-if="btnAuth.takeGoodPeopleInp"
            :disabled="!multipleSelection.length"
            icon="delete"
            style="margin-left: 10px"
            @click="onDelete"
          >
            删除
          </el-button>
        </div>
        <div>
          <div>
            <el-button>
              批量下单
              <input type="file" accept=".xls,.xlsx" @change="readExcel"/>
            </el-button>
          </div>
        </div>
        <div>
          <div>{{amountTitle +":" +totalAmount }}</div>
        </div>
      </div>
      <el-table
        ref="multipleTable"
        :data="tableData"
        row-key="ID"
        style="width: 100%"
        tooltip-effect="dark"
        @selection-change="handleSelectionChange"
      >
        <el-table-column
          v-if="btnAuth.takeGoodPeopleInp"
          type="selection"
          width="55"
        />
        <el-table-column
          align="left"
          label="单据编号"
          prop="billNumber"
          width="280"
        />
        <el-table-column align="left" label="下单日期" width="180">
          <template #default="scope">
            {{ formatDate(scope.row.CreatedAt) }}
          </template>
        </el-table-column>

        <el-table-column
          align="left"
          label="下单人"
          prop="createdBySimpleUser.nickName"
          width="120"
        />
        // 下单设备
        <el-table-column
          align="left"
          label="下单设备"
          prop="device"
          width="120"
        >
          <template #default="scope">
            <div>
              {{ scope.row.device == 0 ? "小程序" : "网页端" }}
            </div>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="档口名"
          prop="stall.stall"
          width="120"
        />
        <el-table-column
          align="left"
          label="档口号"
          prop="stall.stallNumber"
          width="120"
        />
        <el-table-column
          align="left"
          label="拿货人"
          prop="takeGoodPeople.nickName"
          width="180"
        >
          <template #default="scope">
            <div>
              <el-select
                v-if="btnAuth.takeGoodPeopleInp"
                v-model="scope.row.takeGoodPeopleId"
                filterable
                placeholder="请选择司机"
                @change="takeGoodPeopleInpChange(scope.row)"
              >
                <el-option
                  v-for="(item, index) in userOption"
                  :key="index"
                  :label="item.nickName"
                  :value="item.ID"
                />
              </el-select>
              <span v-else>{{ scope.row.takeGoodPeople.nickName }}</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="市场"
          prop="stall.market.marketName"
          width="120"
        />
        <el-table-column align="left" label="加急" prop="urgent" width="120">
          <template #default="scope">
            <div>
              {{ scope.row.urgent ? "是" : "否" }}
            </div>
          </template>
        </el-table-column>
        <!-- <el-table-column
          align="left"
          label="档口价格"
          prop="stall.price"
          width="120"
        /> -->
        <el-table-column
          align="left"
          label="折扣后金额"
          prop="discountAmount"
          width="120"
        />
        <!-- <el-table-column
          align="left"
          label="折扣率"
          prop="discountRate"
          width="120"
        /> -->
        <el-table-column
          align="left"
          label="拿货数量"
          prop="takeGoodNum"
          width="120"
        />
        <el-table-column align="left" label="备注" prop="remarks" width="120"  :show-overflow-tooltip='true'/>
        <el-table-column
          align="left"
          label="下单人留言"
          prop="orderMessage"
          width="120"
        />
        <el-table-column
          v-if="btnAuth.takeGoodPeopleInp"
          align="left"
          label="管理员留言"
          prop="adminMessage"
          width="120"
        />
        <el-table-column
          align="left"
          label="司机留言"
          prop="driverMessage"
          width="120"
        />
        <el-table-column align="left" label="货单状态" width="180">
          <template #default="scope">
            <div>
              <el-select
                v-if="btnAuth.takeGoodPeopleInp"
                v-model="scope.row.goodBillStatus"
                filterable
                placeholder="请选货单状态"
                value-key="ID"
                @change="goodBillStatusChange(scope.row)"
              >
                <el-option
                  v-for="(item, index) in goodBillStatusOption"
                  :key="index"
                  :label="item.name"
                  :value="item"
                />
              </el-select>
              <span v-else>{{ scope.row.goodBillStatus.name }}</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="完成状态"
          prop="finishStatus"
          width="120"
        >
          <template #default="scope">
            <el-switch
              v-model="scope.row.finishStatus"
              :disabled="!btnAuth.takeGoodPeopleInp"
              @change="finishStatusChange(scope.row)"
            />
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="司机核实"
          prop="driverVerify"
          width="120"
        >
          <template #default="scope">
            <el-switch
              v-model="scope.row.driverVerify"
              :disabled="!btnAuth.takeGoodPeopleInp"
              @change="finishStatusChange(scope.row)"
            />
          </template>
        </el-table-column>
        <el-table-column label="完成时间" prop="finishTime" width="180" />
        <el-table-column
          label="完成人"
          prop="finishPeople.nickName"
          width="180"
        />
        <el-table-column label="支付方式" prop="payType" width="180" />
        <el-table-column label="是否需要审核" prop="isManual" width="180">
          <template #default="scope">
            <div>
              {{ scope.row.isManual=='0' ?  "不需要" : "需要" }}
            </div>
          </template>
        </el-table-column>
        <el-table-column label="微信订单号" prop="wechatOrderId" width="280">
          <template #default="scope">
            <div >
              {{ scope.row.wechatOrderId === "" ? "无" : scope.row.wechatOrderId }}
            </div>
          </template>
        </el-table-column>
        <el-table-column label="是否支付完成" prop="isPay" width="180">
          <template #default="scope">
            <div>
              {{ scope.row.isPay==='1'? "已支付" : "未支付" }}
            </div>
          </template>
        </el-table-column>
        <el-table-column label="售后状态" prop="afterSaleStatus" width="180">
          <template #default="scope">
            <div>
              {{ scope.row.afterSaleStatus === '0' ? "无售后": '退款中'}}
            </div>
          </template>
        </el-table-column>
        <el-table-column label="同意退款" prop="agreeRefund" width="180">
          <template #default="scope">
            <div v-if="scope.row.afterSaleStatus=== '0' ">
              {{ "无售后" }}
            </div>
            <div v-if="scope.row.afterSaleStatus=== '1'">
              {{ scope.row.agreeRefund === "0" ? "未处理" : (scope.row.agreeRefund === 1 ? "同意" : "拒绝") }}
            </div>
          </template>
        </el-table-column>
        <el-table-column label="退款完成" prop="refundStatus" width="180">
          <template #default="scope">
            <div v-if="scope.row.afterSaleStatus=== '0' ">
              {{ "无售后" }}
            </div>
            <div v-if="scope.row.afterSaleStatus=== '1'">
              {{ scope.row.refundStatus === '0'? '未完成':'已完成' }}
            </div>
          </template>
        </el-table-column>
        <el-table-column label="退款订单号" prop="refundOrderId" width="180">
          <template #default="scope">
            <div v-if="scope.row.afterSaleStatus=== '0' ">
              {{ "无售后" }}
            </div>
            <div v-if="scope.row.afterSaleStatus=== '1'">
              {{ scope.row.refundOrderId === ''? '未完成' : scope.row.refundOrderId }}
            </div>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          fixed="right"
          label="操作"
          min-width="180"
        >
          <template #default="scope">
            <el-button
              class="table-button"
              icon="edit"
              link
              type="primary"
              @click="updateGoodBillFunc(scope.row)"
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
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100, 500, 1000, 2000]"
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
        <el-form-item label="加急:" prop="urgent">
          <el-switch v-model="formData.urgent" @change="urgentChange" />
        </el-form-item>
        <el-form-item
          :label="formData.urgent ? '加急档口' : '档口'"
          prop="stall"
        >
          <el-select
            v-model="formData.stall"
            filterable
            :placeholder="formData.urgent ? '请选择加急档口' : '请选择档口'"
            value-key="ID"
          >
            <el-option
              v-for="stall in stallOptions"
              :key="stall.ID"
              :label="stall.stall"
              :value="stall"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="拿货数量:" prop="takeGoodNum">
          <el-input-number
            v-model="formData.takeGoodNum"
            :clearable="true"
            placeholder="请输入拿货数量"
            style="width: 200px"
          />
        </el-form-item>
        <el-form-item label="下单人留言:" prop="orderMessage">
          <el-input
            v-model="formData.orderMessage"
            :clearable="true"
            placeholder="请输入下单人留言"
          />
        </el-form-item>
        <el-form-item
          v-if="btnAuth.takeGoodPeopleInp"
          label="管理员留言:"
          prop="adminMessage"
        >
          <el-input
            v-model="formData.adminMessage"
            :clearable="true"
            placeholder="请输入管理员留言"
          />
        </el-form-item>
        <el-form-item
          v-if="btnAuth.takeGoodPeopleInp"
          label="司机留言:"
          prop="driverMessage"
        >
          <el-input
            v-model="formData.driverMessage"
            :clearable="true"
            placeholder="请输入司机留言"
          />
        </el-form-item>
        <el-form-item label="收件地址:" prop="remarks">
          <el-input
            v-model="formData.remarks"
            :clearable="true"
            placeholder="请输入备注"
          />
        </el-form-item>
        <upload-image
          v-if="btnAuth.uploadImage"
          button-name="上传图片"
          @on-success="getImageList"
        />
        <div
          style="
            display: grid;
            grid-template-columns: repeat(auto-fill, minmax(110px, 1fr));
            margin-top: 10px;
          "
        >
          <div
            v-for="img in formData.images"
            :key="img"
            class="imgItem"
            style="position: relative; width: 100px; height: 100px"
          >
            <el-icon
              class="delete-icon"
              style="
                position: absolute;
                top: 50%;
                left: 50%;
                transform: translate(-50%, -50%);
                z-index: 2;
                font-size: 20px;
              "
              @click="deleteImage(img)"
            >
              <Delete />
            </el-icon>
            <el-image
              :src="img"
              fit="cover"
              style="width: 100px; height: 100px"
            />
          </div>
        </div>
      </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createGoodBill,
  deleteGoodBill,
  deleteGoodBillByIds,
  findGoodBill,
  getGoodBillList,
  updateGoodBill,
} from "@/api/bill/goodBill";
import { getGoodBillStatusList } from "@/api/dataConfig/goodBillStatus";
// 全量引入格式化工具 请按需保留
import { formatDate } from "@/utils/format";

import { ElMessage, ElMessageBox } from "element-plus";
import { computed, onMounted, reactive, ref } from "vue";
import { getStallList } from "@/api/dataConfig/stall";
import { getUserList } from "@/api/user";

import { useBtnAuth } from "@/utils/btnAuth";
import UploadImage from "@/components/upload/image.vue";

import * as XLSX from 'xlsx'
import { getMarketList } from "@/api/dataConfig/market";
import { screen } from "@/utils/screen";

const btnAuth = useBtnAuth();

defineOptions({
  name: "GoodBill",
});

const getImageList = (url) => {
  formData.value.images.push(`api/${url}`);
};

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  remarks: "",
  stall: null,
  urgent: false,
  images: [],
});

// 验证规则
const rule = reactive({
  stall: [{ required: true, message: "请选择档口", trigger: "change" }],
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
const amountTitle=ref('合计金额')

const stallOptions = ref([]);

// ------批量下单部分
// 批量下单之读取Excel
const readExcel = (event) => {
  const fileInput = event.target;
  if (fileInput.files.length > 0) {
    const file = fileInput.files[0];

    // 获取文件的扩展名
    const fileName = file.name;
    const fileExtension = fileName.split('.').pop().toLowerCase();

    // 检查文件扩展名是否为 .xls 或 .xlsx
    if (fileExtension === 'xls' || fileExtension === 'xlsx') {
      const reader = new FileReader();
      reader.onload = async (e) => {
        const data = new Uint8Array(e.target.result);
        const workbook = XLSX.read(data, { type: 'array' });
        const worksheet = workbook.Sheets[workbook.SheetNames[0]];
        const fileList = XLSX.utils.sheet_to_json(worksheet, { header: 1 });
        fileList.shift();
        
        // 处理解析后的数据
        for (let i = 0; i < fileList.length; i++) {
          // 解析后的数据fileList，fileList[i][1]是档口名，fileList[i][2]是档口编号
          if (fileList[i].length > 0) {
            screen(fileList[i], fileList[i][1], fileList[i][2]);
            getTableData();
          }
        }

        // 逻辑处理完后清除文件
        fileInput.value = '';  // 清除文件
      };
      reader.readAsArrayBuffer(file);
    } else {
      alert('请上传正确的Excel表格');
      // 清除无效的文件
      fileInput.value = '';
    }
  }
};



// 合计金额
const totalAmount = computed(() => {
  return tableData.value.reduce((total, item) => {
    return total + item.stall.price;
  }, 0);
});

const queryStallList = async () => {
  const table = await getStallList({
    page: page.value,
    pageSize: pageSize.value,
    urgent: formData.value.urgent,
  });
  stallOptions.value = table.data.list;
};
queryStallList();

const userOption = ref([]);

const queryUserList = async () => {
  const res = await getUserList({
    page: 1,
    pageSize: 1000,
    authorityIds: [999],
  });
  userOption.value = res.data.list;
};

const createdOption = ref([]);

const querycreatedList = async () => {
  const res = await getUserList({
    page: 1,
    pageSize: 1000,
    authorityIds: [998],
  });
  createdOption.value = res.data.list;
};

if (btnAuth.takeGoodPeopleInp) {
  queryUserList();
  querycreatedList();
}
const goodBillStatusOption = ref([]);
const queryGoodBillStatus = async () => {
  const res = await getGoodBillStatusList({
    page: 1,
    pageSize: 1000,
  });
  goodBillStatusOption.value = res.data.list;
};
queryGoodBillStatus();
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
  const table = await getGoodBillList({
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
  dataChangeIng.value = true;
  setTimeout(() => {
    dataChangeIng.value = false;
  }, 1000);
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
  console.log(row.CreatedAt, "row");
  // 计算下单时间有没有超过半小时
  const time = new Date().getTime() - new Date(row.CreatedAt).getTime();
  if (time > 1000 * 60 * 30 && !btnAuth.takeGoodPeopleInp) {
    ElMessage({
      type: "warning",
      message: "下单时间超过半小时，不可删除，请联系管理员",
      duration: 3000,
    });
    return;
  }
  ElMessageBox.confirm("确定要删除吗?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  }).then(() => {
    deleteGoodBillFunc(row);
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
    const res = await deleteGoodBillByIds({ IDs });
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
const updateGoodBillFunc = async (row) => {
  const res = await findGoodBill({ ID: row.ID });
  type.value = "update";
  if (res.code === 0) {
    formData.value = res.data.regb;
    formData.value.images = !formData.value.images
      ? []
      : JSON.parse(formData.value.images);
    dialogFormVisible.value = true;
  }
};

// 删除行
const deleteGoodBillFunc = async (row) => {
  const res = await deleteGoodBill({ ID: row.ID });
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
    remarks: "",
    stall: null,
    images: [],
  };
};

const dataChangeIng = ref(false);
// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return;
    let res;
    formData.value.images = JSON.stringify(formData.value.images);
    switch (type.value) {
      case "create":
        res = await createGoodBill(formData.value);
        break;
      case "update":
        res = await updateGoodBill({
          ...formData.value,
          stallId: formData.value.stall.ID,
          isManual:'0'
        });
        break;
      default:
        res = await createGoodBill(formData.value);
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

const takeGoodPeopleInpChange = async (val) => {
  await updateGoodBill(val);
  ElMessage({
    type: "success",
    message: "拿货人修改成功",
  });
};
const urgentChange = () => queryStallList();

const finishStatusChange = async (val) => {
  if (dataChangeIng.value) return;
  await updateGoodBill(val);
  ElMessage({
    type: "success",
    message: "状态修改成功",
  });
  onSubmit();
};

const goodBillStatusChange = async (val) => {
  await updateGoodBill({
    ...val,
    goodBillStatusId: val.goodBillStatus.ID,
  });
  ElMessage({
    type: "success",
    message: "货单状态修改成功",
  });
  onSubmit();
};

const deleteImage = (img) => {
  formData.value.images = formData.value.images.filter((item) => item !== img);
};
</script>

<style>
.delete-icon {
  display: none;
}

.imgItem:hover .delete-icon {
  display: block;
}
</style>
