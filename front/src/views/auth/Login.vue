<template>
  <div id="login">
    <h1>登录</h1>
    <el-form
        ref="ruleFormRef"
        :model="loginForm"
        status-icon
        :rules="rules"
        label-width="120px"
    >
      <el-form-item label="Phone" prop="phone">
        <el-input v-model="loginForm.phone" type="password" autocomplete="off"/>
      </el-form-item>
      <el-form-item label="Password" prop="password">
        <el-input v-model="loginForm.password" type="password" autocomplete="off"/>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitForm(ruleFormRef)"
        >Submit
        </el-button
        >
        <el-button @click="resetForm(ruleFormRef)">Reset</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>


<script setup lang="ts">
import {reactive, ref} from "vue"
import {useRouter} from "vue-router";
import {login} from "@/api/login"
import {ElMessage, type FormInstance} from 'element-plus'
const router = useRouter()
//#region loginForm表单校验器
const ruleFormRef = ref<FormInstance>()
const rules = reactive({
  phone: [{required: true, trigger: 'blur'}, {min: 6, message: 'Length should be 6', trigger: 'blur'}],
  password: [{required: true, trigger: 'blur'}, {min: 6, message: 'Length should be 6', trigger: 'blur'}],
})
//#endregion 校验器

const submitForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.validate(async (valid) => {
    if (valid) {
      const resp = await login(loginForm)
      await window.localStorage.setItem("custom-token", "Bearer "+<string>resp.data?.token)
      await router.push('/home')
    } else {
      await ElMessage.error("请正确填写表单")
      return false
    }
  })
}

const resetForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.resetFields()
}

const loginForm = reactive({
  phone: '',
  password: '',
  role: 1,
})
</script>

<style lang="scss" scoped>
  #login{
    width: 100%;
    display: flex;
    flex-direction: column;
    align-items: center;
    h1{
      margin: 20px;
    }
  }
</style>