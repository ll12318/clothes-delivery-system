<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="档口名称:" prop="stall">
          <el-input v-model="formData.stall" :clearable="true"  placeholder="请输入档口名称" />
       </el-form-item>
        <el-form-item label="档口号:" prop="stallNumber">
          <el-input v-model="formData.stallNumber" :clearable="true"  placeholder="请输入档口号" />
       </el-form-item>
        <el-form-item label="备注:" prop="remarks">
          <el-input v-model="formData.remarks" :clearable="true"  placeholder="请输入备注" />
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
  createStall,
  updateStall,
  findStall
} from '@/api/dataConfig/stall'

defineOptions({
    name: 'StallForm'
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
            stall: '',
            stallNumber: '',
            remarks: '',
        })
// 验证规则
const rule = reactive({
               stall : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               stallNumber : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findStall({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.restall
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
               res = await createStall(formData.value)
               break
             case 'update':
               res = await updateStall(formData.value)
               break
             default:
               res = await createStall(formData.value)
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
