

// 筛选手动和自动订单
let formData={
  isManual:'0', // 0自动  1手动
  takeGoodNum:0,
  stall:{},
  urgent:false,
  images:[],
  orderMessage	:''
}
import { createGoodBill } from "@/api/bill/goodBill"
import { getStallList } from "@/api/dataConfig/stall"

export const screen = async(fileList,stall,stallNumber)=>{
  let urgent = fileList[5]==='是'? true : false
  let  stallList = await getStallList({
    page:1,
    pageSize:100,
    stall,
    stallNumber,
    urgent
  })
  compare(fileList,stallList,urgent)
}
const compare = async (fileList, stallList,urgent) => {
  formData.orderMessage	=''
  let found = false;
  formData.images=[]
  for (let i = 0; i < stallList.data.list.length; i++) {
    if (stallList.data.list.length !== 0) {
      if (stallList.data.list[i].stall === fileList[1]) {
        found = true;
        formData.stall = stallList.data.list[i]
        formData.takeGoodNum = fileList[3]
        formData.urgent = urgent
        formData.isManual = '0'
        formData.images = JSON.stringify(formData.images);
        createGoodBill(formData)
        break; // 找到后跳出循环
      }
    }
  }

  if (!found) {
    // formData.images=[]
    // 如果没有找到，可以在这里调用 getStallList
    let nstall = await getStallList({
      page: 1,
      pageSize: 100,
      stall:'审核中',
      stallNumber:'审核中',
    })
    formData.stall =nstall.data.list[0]
    formData.images = JSON.stringify([]);
    formData.takeGoodNum = fileList[3]
    formData.isManual = '1'
    formData.orderMessage	 = `市场名：${fileList[0]}，档口名：${fileList[1]}，档口编号：${fileList[2]}，是否加急：${fileList[5]}，拿货件数：${fileList[3]}，报单人：${fileList[4]}`
    createGoodBill(formData)

  }
};


// 获取审核档口
