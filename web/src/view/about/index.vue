<template>
  <el-switch v-model="finishStatus" @change="finishStatusChange"></el-switch>
  <el-tabs type="border-card" @tab-click="tabClcik">
    <el-tab-pane v-for="mk in marketList" :key="mk.ID" :label="mk.marketName"
      >User</el-tab-pane
    >
  </el-tabs>
</template>

<script setup>
// getGoodBillListByDriver
import { ref } from "vue";
import {
  getGoodBillMarketListByDriver,
  getGoodBillListByMarketId,
} from "@/api/bill/goodBill";

const finishStatus = ref(false);

const finishStatusChange = (val) => {
  queryGoodBillListByDriver();
};

const tabClcik = async (tab) => {
  const res = await getGoodBillListByMarketId({
    finishStatus: finishStatus.value,
    marketId: marketList.value[tab.index].ID,
  });
};

const marketList = ref([]);
const queryGoodBillListByDriver = async () => {
  const res = await getGoodBillMarketListByDriver({
    finishStatus: finishStatus.value,
  });
  marketList.value = res.data.list;
};
queryGoodBillListByDriver();
</script>

<style scoped></style>
