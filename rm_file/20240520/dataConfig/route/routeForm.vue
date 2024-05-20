<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="路线名称:" prop="routeName">
          <el-input v-model="formData.routeName" :clearable="true"  placeholder="请输入路线名称" />
       </el-form-item>
        <el-form-item label="备注:" prop="remarks">
          <el-input v-model="formData.remarks" :clearable="true"  placeholder="请输入备注" />
       </el-form-item>
        <el-form-item label="档口:" prop="stalls">
        <el-select  multiple  v-model="formData.stalls" placeholder="请选择档口" style="width:100%" :clearable="true" >
          <el-option v-for="(item,key) in dataSource.stalls" :key="key" :label="item.label" :value="item.value" />
        </el-select>
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
    getRouteDataSource,
  createRoute,
  updateRoute,
  findRoute
} from '@/api/dataConfig/route'

defineOptions({
    name: 'RouteForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            routeName: '',
            remarks: '',
            stalls: [],
        })
// 验证规则
const rule = reactive({
               routeName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getRouteDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findRoute({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reroute
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createRoute(formData.value)
               break
             case 'update':
               res = await updateRoute(formData.value)
               break
             default:
               res = await createRoute(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
