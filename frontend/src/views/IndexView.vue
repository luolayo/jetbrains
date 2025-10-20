<template>
  <div class="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100 flex items-center justify-center p-4">
    <!-- 加载验证中 -->
    <div v-if="status === 'loading'" class="bg-white rounded-2xl shadow-xl p-8 max-w-md w-full text-center">
      <div class="mb-6">
        <div class="inline-flex items-center justify-center w-20 h-20 bg-blue-100 rounded-full mb-4">
          <svg class="animate-spin h-10 w-10 text-blue-600" xmlns="http://www.w3.org/2000/svg" fill="none"
               viewBox="0 0 24 24">
            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
            <path class="opacity-75" fill="currentColor"
                  d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
          </svg>
        </div>
      </div>
      <h2 class="text-2xl font-bold text-gray-800 mb-2">正在验证设备</h2>
      <p class="text-gray-500">请稍候，系统正在验证您的访问权限...</p>
    </div>

    <!-- 需要输入订单号 -->
    <div v-else-if="status === 'needOrder'" class="bg-white rounded-2xl shadow-xl p-8 max-w-md w-full">
      <div class="text-center mb-6">
        <div class="inline-flex items-center justify-center w-20 h-20 bg-yellow-100 rounded-full mb-4">
          <svg class="w-10 h-10 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                  d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"></path>
          </svg>
        </div>
        <h2 class="text-2xl font-bold text-gray-800 mb-2">设备未授权</h2>
        <p class="text-gray-500 mb-6">请输入您的订单号以激活设备访问权限</p>
      </div>

      <form @submit.prevent="handleSubmitOrder">
        <div class="mb-6">
          <label class="block text-sm font-medium text-gray-700 mb-2">订单号</label>
          <input
            v-model="orderNumber"
            type="text"
            placeholder="请输入订单号"
            class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent outline-none transition"
            :disabled="verifying"
            required
          />
        </div>

        <div v-if="errorMessage" class="mb-4 p-3 bg-red-50 border border-red-200 rounded-lg">
          <p class="text-sm text-red-600 flex items-center">
            <svg class="w-4 h-4 mr-2" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd"
                    d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z"
                    clip-rule="evenodd"></path>
            </svg>
            {{ errorMessage }}
          </p>
        </div>

        <button
          type="submit"
          :disabled="verifying || !orderNumber"
          class="w-full bg-blue-600 hover:bg-blue-700 disabled:bg-gray-400 text-white font-semibold py-3 px-6 rounded-lg transition duration-200 flex items-center justify-center"
        >
          <svg v-if="verifying" class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg"
               fill="none" viewBox="0 0 24 24">
            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
            <path class="opacity-75" fill="currentColor"
                  d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
          </svg>
          {{ verifying ? '验证中...' : '提交验证' }}
        </button>
      </form>
    </div>

    <!-- 验证成功，进入系统 -->
    <div v-else-if="status === 'success'" class="bg-white rounded-2xl shadow-xl p-8 max-w-4xl w-full">
      <div class="text-center mb-8">
        <div class="inline-flex items-center justify-center w-20 h-20 bg-green-100 rounded-full mb-4">
          <svg class="w-10 h-10 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                  d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
          </svg>
        </div>
        <h2 class="text-3xl font-bold text-gray-800 mb-2 text-center">验证成功！</h2>
        <p class="text-gray-500 text-center">欢迎进入系统</p>
        <p class="text-gray-500 text-center">闲鱼关注：养小猫_ 其余均为倒卖狗，倒卖死全家</p>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8">
        <div class="bg-gradient-to-br from-blue-50 to-blue-100 rounded-xl p-6 border border-blue-200">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-800">设备信息</h3>
            <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                    d="M9.75 17L9 20l-1 1h8l-1-1-.75-3M3 13h18M5 17h14a2 2 0 002-2V5a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"></path>
            </svg>
          </div>
          <p class="text-sm text-gray-600">设备已授权</p>
          <p class="text-xs text-gray-500 mt-1">ID: {{ drvice.drviceInfo.uuid }}</p>
        </div>

        <div class="bg-gradient-to-br from-green-50 to-green-100 rounded-xl p-6 border border-green-200">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-800">订单状态</h3>
            <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                    d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
          <p class="text-sm text-gray-600">已激活</p>
          <p class="text-xs text-gray-500 mt-1">订单号: {{ orderNumber || 'AUTO' }}</p>
        </div>

        <div class="bg-gradient-to-br from-purple-50 to-purple-100 rounded-xl p-6 border border-purple-200">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-800">访问权限</h3>
            <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                    d="M15 7a2 2 0 012 2m4 0a6 6 0 01-7.743 5.743L11 17H9v2H7v2H4a1 1 0 01-1-1v-2.586a1 1 0 01.293-.707l5.964-5.964A6 6 0 1121 9z"></path>
            </svg>
          </div>
          <p class="text-sm text-gray-600">完全访问</p>
          <p class="text-xs text-gray-500 mt-1">有效期: 永久</p>
        </div>
      </div>

      <div class="space-y-3">
        <button
          @click="$router.push('/download')"
          class="w-full bg-blue-600 hover:bg-blue-700 text-white font-semibold py-4 rounded-lg transition duration-200 flex items-center justify-center group"
        >
          <svg class="w-5 h-5 mr-2 group-hover:animate-bounce" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                  d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4"/>
          </svg>
          下载软件(暂不可用，等待更新)
        </button>

        <button
          @click="$router.push('/activation')"
          class="w-full bg-indigo-600 hover:bg-indigo-700 text-white font-semibold py-4 rounded-lg transition duration-200 flex items-center justify-center group"
        >
          <svg class="w-5 h-5 mr-2 group-hover:scale-110 transition-transform" fill="none" stroke="currentColor"
               viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                  d="M15 7a2 2 0 012 2m4 0a6 6 0 01-7.743 5.743L11 17H9v2H7v2H4a1 1 0 01-1-1v-2.586a1 1 0 01.293-.707l5.964-5.964A6 6 0 1121 9z"/>
          </svg>
          激活软件
        </button>

        <button
          @click="clean"
          class="w-full bg-indigo-600 hover:bg-indigo-700 text-white font-semibold py-4 rounded-lg transition duration-200 flex items-center justify-center group"
        >
          <svg class="w-5 h-5 mr-2 group-hover:scale-110 transition-transform" fill="none" stroke="currentColor"
               viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                  d="M15 7a2 2 0 012 2m4 0a6 6 0 01-7.743 5.743L11 17H9v2H7v2H4a1 1 0 01-1-1v-2.586a1 1 0 01.293-.707l5.964-5.964A6 6 0 1121 9z"/>
          </svg>
          清理激活
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ElMessage } from 'element-plus'
import {ref, onMounted} from 'vue'
import {useDrvice} from "../store/useDrvice";
import {decrypt} from "../util/crypt";
import {Verification} from "../api";
import {Clean} from "../../wailsjs/go/main/App";

type VerificationStatus = 'loading' | 'needOrder' | 'success'


const drvice = useDrvice()

const status = ref<VerificationStatus>('loading')
const orderNumber = ref('')
const verifying = ref(false)
const errorMessage = ref('')


const initialVerification = async (): Promise<{ data: string }> => {
  return Verification(drvice.drviceInfo.uuid, drvice.drviceInfo.mac)
}

const verifyOrderNumber = async (order: string): Promise<{data: string}> => {
  const res = await fetch("http://localhost:8080/api/order/device/add", {
    method: "POST",
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      orderTid: order,
      uuid: drvice.drviceInfo.uuid,
      mac: drvice.drviceInfo.mac,
    })
  })
  return res.json()
}

// 页面加载时执行初始验证
onMounted(async () => {

  const result = await initialVerification()
  const dataString = await decrypt(result.data, drvice.code)
  const data = JSON.parse(dataString)
  if (data.code === 200) {
    status.value = 'success'
    orderNumber.value = data.msg
  } else if (result) {
    ElMessage.error('设备未授权，请输入订单号激活')
    status.value = 'needOrder'
  }
})

// 处理订单号提交
const handleSubmitOrder = async () => {
  if (!orderNumber.value.trim()) {
    errorMessage.value = '请输入订单号'
    return
  }

  errorMessage.value = ''
  verifying.value = true

  const result = await verifyOrderNumber(orderNumber.value)
  const dataString = await decrypt(result.data, drvice.code)
  const data = JSON.parse(dataString)
  console.log(data)
  if (data.code !== 200){
    verifying.value = false
    ElMessage.error(data.msg || '验证失败，请重试')
    return
  }
    status.value = 'success'
  ElMessage({
    message: '绑定成功，正在进入系统...',
    type: 'success',
  })
}

const clean = async () => {
  const data = await Clean()
  ElMessage({
    message: data,
    type: 'success',
  })
}

</script>