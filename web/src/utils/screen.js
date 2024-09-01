// 筛选手动和自动订单
let formData = {
  isManual: '0', // 0自动  1手动
  takeGoodNum: 0,
  stall: {},
  urgent: false,
  images: [],
  orderMessage: ''
};
// 清理字符串函数
export const cleanString = (str) => {
  if (typeof str !== 'string') {
    return '';  // 如果传入的不是字符串，返回一个空字符串或适当的默认值
  }
  
  return str.replace(/^\s+|\s+$/g, '')      // 清除前后空格
            .replace(/[\r\n]+/g, '')        // 清除换行符
            .replace(/\u200B/g, '')         // 清除零宽字符
            .replace(/[\x00-\x1F\x7F]/g, '') // 清除控制字符
            .replace(/\s+/g, '');           // 清除所有空格
};

import { createGoodBill } from "@/api/bill/goodBill";
import { getStallList } from "@/api/dataConfig/stall";

export const screen = async (fileList, stall, stallNumber) => {
  let urgent = fileList[5] === '是' ? true : false;
  let stallList = await getStallList({
    page: 1,
    pageSize: 100,
    stall,
    stallNumber,
    urgent
  });
  compare(fileList, stallList, urgent,stall);
};

const compare = async (fileList, stallList, urgent,stall) => {
  formData.orderMessage = '';
  let found = false;
  formData.images = [];

  for (let i = 0; i < stallList.data.list.length; i++) {
    if (stallList.data.list.length !== 0) {
      if (stallList.data.list[i].stall === stall) {
        found = true;
        formData.stall = stallList.data.list[i];
        formData.takeGoodNum = fileList[3];
        formData.urgent = urgent;
        formData.isManual = '0';
        formData.images = JSON.stringify(formData.images);
        formData.orderMessage = fileList[6];
        formData.declarant = String(fileList[4]);
        createGoodBill(formData);
        break; // 找到后跳出循环
      }
    }
  }

  if (!found) {
    let nstall = await getStallList({
      page: 1,
      pageSize: 100,
      stall: '审核中',
      stallNumber: '审核中',
    });
    formData.stall = nstall.data.list[0];
    formData.images = JSON.stringify([]);
    formData.takeGoodNum = fileList[3];
    formData.isManual = '1';
    formData.declarant = String(fileList[4]);
    formData.orderMessage = `市场名：${fileList[0]}，档口名：${fileList[1]}，档口编号：${fileList[2]}，是否加急：${fileList[5]}，拿货件数：${fileList[3]}，报单人：${fileList[4]},备注：${fileList[6]}`;
    createGoodBill(formData);
  }
};