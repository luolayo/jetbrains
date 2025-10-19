<template>
  <div class="min-h-screen bg-gradient-to-br from-slate-50 to-slate-100">
    <div class="container mx-auto px-4 py-8 max-w-4xl">
      <!-- 返回按钮 -->
      <button
        @click="goBack"
        class="mb-6 flex items-center gap-2 text-slate-600 hover:text-slate-900 transition-colors"
      >
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18" />
        </svg>
        返回首页
      </button>

      <!-- 页面标题 -->
      <div class="text-center mb-8">
        <h1 class="text-4xl font-bold text-slate-900 mb-2">激活软件</h1>
        <p class="text-slate-600">一键激活您的 JetBrains 产品</p>
      </div>

      <!-- 长提示卡片 -->
      <div class="bg-gradient-to-r from-blue-50 to-indigo-50 border-2 border-blue-200 rounded-2xl p-6 mb-6">
        <div class="flex gap-4">
          <div class="flex-shrink-0">
            <div class="w-12 h-12 bg-blue-600 rounded-xl flex items-center justify-center">
              <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
            </div>
          </div>
          <div class="flex-1">
            <h3 class="text-lg font-semibold text-slate-900 mb-2">激活须知</h3>
            <ul class="space-y-2 text-sm text-slate-700">
              <li class="flex items-start gap-2">
                <svg class="w-5 h-5 text-blue-600 flex-shrink-0 mt-0.5" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
                </svg>
                <span>请确保您已经下载并安装了需要激活的软件</span>
              </li>
              <li class="flex items-start gap-2">
                <svg class="w-5 h-5 text-blue-600 flex-shrink-0 mt-0.5" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
                </svg>
                <span>激活过程需要网络连接，请保持网络畅通</span>
              </li>
              <li class="flex items-start gap-2">
                <svg class="w-5 h-5 text-blue-600 flex-shrink-0 mt-0.5" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
                </svg>
                <span>激活成功后，需要手动输入激活码，激活码会在激活信息中告诉您</span>
              </li>
              <li class="flex items-start gap-2">
                <svg class="w-5 h-5 text-blue-600 flex-shrink-0 mt-0.5" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
                </svg>
                <span>如遇到问题，请联系卖家</span>
              </li>
            </ul>
          </div>
        </div>
      </div>

      <!-- 版本选择区域 -->
      <div class="bg-white rounded-2xl shadow-sm p-6 mb-6">
        <h2 class="text-lg font-semibold text-slate-900 mb-4">选择版本类型</h2>

        <!-- 提示信息 -->
        <div class="bg-amber-50 border-2 border-amber-200 rounded-xl p-4 mb-4">
          <div class="flex gap-3">
            <svg class="w-5 h-5 text-amber-600 flex-shrink-0 mt-0.5" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd" d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z" clip-rule="evenodd" />
            </svg>
            <div class="flex-1">
              <p class="text-sm font-medium text-amber-900 mb-1">版本选择说明</p>
              <p class="text-sm text-amber-800">
                如果您电脑安装了低版本请选择低版本，如果您电脑全是最新版请选择高版本。<br>
                <strong>2024年及以后版本为高版本，2024年以前版本为低版本。</strong>
              </p>
            </div>
          </div>
        </div>

        <!-- 版本选项 -->
        <div class="grid md:grid-cols-2 gap-4">
          <button
            @click="selectedVersion = 'high'"
            :class="[
              'p-5 rounded-xl border-2 transition-all text-left',
              selectedVersion === 'high'
                ? 'border-blue-500 bg-blue-50 shadow-md'
                : 'border-slate-200 hover:border-slate-300 hover:shadow-sm'
            ]"
          >
            <div class="flex items-start gap-3">
              <div :class="[
                'w-6 h-6 rounded-full border-2 flex items-center justify-center flex-shrink-0 mt-0.5',
                selectedVersion === 'high' ? 'border-blue-500 bg-blue-500' : 'border-slate-300'
              ]">
                <div v-if="selectedVersion === 'high'" class="w-2.5 h-2.5 bg-white rounded-full"></div>
              </div>
              <div class="flex-1">
                <h3 class="font-semibold text-slate-900 mb-1">高版本</h3>
                <p class="text-sm text-slate-600">适用于 2024 年及以后发布的版本</p>
                <div class="mt-2 flex items-center gap-1.5">
                  <span class="text-xs bg-green-100 text-green-700 px-2 py-0.5 rounded-full font-medium">
                    2024+
                  </span>
                </div>
              </div>
            </div>
          </button>

          <button
            @click="selectedVersion = 'low'"
            :class="[
              'p-5 rounded-xl border-2 transition-all text-left',
              selectedVersion === 'low'
                ? 'border-blue-500 bg-blue-50 shadow-md'
                : 'border-slate-200 hover:border-slate-300 hover:shadow-sm'
            ]"
          >
            <div class="flex items-start gap-3">
              <div :class="[
                'w-6 h-6 rounded-full border-2 flex items-center justify-center flex-shrink-0 mt-0.5',
                selectedVersion === 'low' ? 'border-blue-500 bg-blue-500' : 'border-slate-300'
              ]">
                <div v-if="selectedVersion === 'low'" class="w-2.5 h-2.5 bg-white rounded-full"></div>
              </div>
              <div class="flex-1">
                <h3 class="font-semibold text-slate-900 mb-1">低版本</h3>
                <p class="text-sm text-slate-600">适用于 2024 年以前发布的版本</p>
                <div class="mt-2 flex items-center gap-1.5">
                  <span class="text-xs bg-slate-100 text-slate-700 px-2 py-0.5 rounded-full font-medium">
                    2023 及更早
                  </span>
                </div>
              </div>
            </div>
          </button>
        </div>
      </div>

      <!-- 激活按钮区域 -->
      <div class="bg-white rounded-2xl shadow-sm p-8 mb-6">
        <div class="text-center">
          <div class="mb-6">
            <div class="w-20 h-20 bg-gradient-to-br from-blue-500 to-indigo-600 rounded-2xl flex items-center justify-center mx-auto mb-4 shadow-lg">
              <svg class="w-10 h-10 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 7a2 2 0 012 2m4 0a6 6 0 01-7.743 5.743L11 17H9v2H7v2H4a1 1 0 01-1-1v-2.586a1 1 0 01.293-.707l5.964-5.964A6 6 0 1121 9z" />
              </svg>
            </div>
            <h2 class="text-2xl font-bold text-slate-900 mb-2">准备激活</h2>
            <p class="text-slate-600">点击下方按钮开始激活您的 JetBrains 产品</p>
          </div>

          <button
            @click="startActivation"
            :disabled="isActivating"
            :class="[
              'px-8 py-4 rounded-xl font-semibold text-lg transition-all shadow-lg',
              isActivating
                ? 'bg-slate-400 cursor-not-allowed'
                : 'bg-gradient-to-r from-blue-600 to-indigo-600 hover:from-blue-700 hover:to-indigo-700 text-white hover:shadow-xl transform hover:scale-105'
            ]"
          >
            <span v-if="!isActivating">立即激活</span>
            <span v-else class="flex items-center justify-center gap-3">
              <svg class="w-6 h-6 animate-spin" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
              {{ activationProgress }}
            </span>
          </button>
        </div>
      </div>

      <!-- 激活信息 -->
      <div class="bg-white rounded-2xl shadow-sm p-6">
        <h2 class="text-lg font-semibold text-slate-900 mb-4">已激活软件</h2>
        <div v-if="activatedSoftwareList.length === 0" class="text-center py-8 text-slate-500">
          <svg class="w-16 h-16 mx-auto mb-3 text-slate-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
          </svg>
          <p>暂无已激活软件</p>
        </div>
        <div v-else class="grid md:grid-cols-2 gap-4">
          <!-- 每个软件一个独立卡片 -->
          <div
            v-for="(software, index) in activatedSoftwareList"
            :key="index"
            :class="[
              'p-4 border-2 rounded-xl transition-all',
              software.status === 'success'
                ? 'border-green-200 bg-green-50 hover:shadow-md'
                : 'border-red-200 bg-red-50'
            ]"
          >
            <div class="flex items-start gap-3">
              <div :class="[
                'w-10 h-10 rounded-lg flex items-center justify-center flex-shrink-0',
                software.status === 'success' ? 'bg-green-100' : 'bg-red-100'
              ]">
                <svg v-if="software.status === 'success'" class="w-6 h-6 text-green-600" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
                </svg>
                <svg v-else class="w-6 h-6 text-red-600" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z" clip-rule="evenodd" />
                </svg>
              </div>
              <div class="flex-1 min-w-0">
                <h3 class="font-semibold text-slate-900 mb-1 truncate">{{ software.productName }}</h3>
                <div class="flex items-center gap-2 mb-1">
                  <span :class="[
                    'px-2 py-0.5 rounded-md text-xs font-semibold',
                    software.status === 'success' ? 'bg-green-600 text-white' : 'bg-red-600 text-white'
                  ]">
                    {{ software.status === 'success' ? '激活成功' : '激活失败' }}
                  </span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 激活结果弹窗 -->
    <transition
      enter-active-class="transition ease-out duration-300"
      enter-from-class="opacity-0"
      enter-to-class="opacity-100"
      leave-active-class="transition ease-in duration-200"
      leave-from-class="opacity-100"
      leave-to-class="opacity-0"
    >
      <div
        v-if="showResultModal"
        class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center p-4 z-50"
        @click="closeModal"
      >
        <transition
          enter-active-class="transition ease-out duration-300"
          enter-from-class="opacity-0 scale-95"
          enter-to-class="opacity-100 scale-100"
          leave-active-class="transition ease-in duration-200"
          leave-from-class="opacity-100 scale-100"
          leave-to-class="opacity-0 scale-95"
        >
          <div
            v-if="showResultModal"
            class="bg-white rounded-2xl shadow-2xl max-w-md w-full p-8"
            @click.stop
          >
            <!-- 成功状态 -->
            <div v-if="activationResult.status === 'success'" class="text-center">
              <div class="w-20 h-20 bg-green-100 rounded-full flex items-center justify-center mx-auto mb-4">
                <svg class="w-10 h-10 text-green-600" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
                </svg>
              </div>
              <h3 class="text-2xl font-bold text-slate-900 mb-2">激活成功！</h3>
              <p class="text-slate-600 mb-6">请将激活码复制到软件里即可成功激活</p>

              <!-- 激活码文本域 -->
              <div class="bg-slate-50 rounded-xl p-4 mb-6 text-left">
                <div class="flex items-center justify-between mb-2">
                  <h4 class="text-sm font-semibold text-slate-900">激活码</h4>
                  <button
                    @click="copyActivationCode"
                    class="text-xs bg-blue-600 text-white px-3 py-1 rounded-md hover:bg-blue-700 transition-colors"
                  >
                    复制
                  </button>
                </div>
                <textarea
                  :value="activationResult.activationCode"
                  readonly
                  class="w-full h-32 px-3 py-2 bg-white border-2 border-slate-200 rounded-lg font-mono text-xs text-slate-700 resize-none focus:outline-none focus:border-blue-500"
                ></textarea>
              </div>

              <button
                @click="closeModal"
                class="w-full py-3 bg-green-600 text-white rounded-xl font-semibold hover:bg-green-700 transition-colors"
              >
                完成
              </button>
            </div>

            <!-- 失败状态 -->
            <div v-else class="text-center">
              <div class="w-20 h-20 bg-red-100 rounded-full flex items-center justify-center mx-auto mb-4">
                <svg class="w-10 h-10 text-red-600" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z" clip-rule="evenodd" />
                </svg>
              </div>
              <h3 class="text-2xl font-bold text-slate-900 mb-2">激活失败</h3>
              <p class="text-slate-600 mb-6">{{ activationResult.errorMessage }}</p>

              <div class="bg-red-50 border-2 border-red-200 rounded-xl p-4 mb-6 text-left">
                <h4 class="font-semibold text-red-900 mb-2">唯一的原因：</h4>
                <ul class="space-y-1 text-sm text-red-700">
                  <li>软件没有安装</li>
                </ul>
              </div>

              <div class="flex gap-3">
                <button
                  @click="closeModal"
                  class="flex-1 py-3 bg-slate-200 text-slate-700 rounded-xl font-semibold hover:bg-slate-300 transition-colors"
                >
                  取消
                </button>
                <button
                  @click="retryActivation"
                  class="flex-1 py-3 bg-blue-600 text-white rounded-xl font-semibold hover:bg-blue-700 transition-colors"
                >
                  重试
                </button>
              </div>
            </div>
          </div>
        </transition>
      </div>
    </transition>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

interface ActivatedSoftware {
  productName: string
  time: string
  status: 'success' | 'failed'
}

interface ActivationResult {
  status: 'success' | 'failed'
  activationCode?: string
  errorMessage?: string
}

// 状态
const selectedVersion = ref<'high' | 'low'>('high')
const isActivating = ref(false)
const activationProgress = ref('正在连接服务器...')
const showResultModal = ref(false)
const activatedSoftwareList = ref<ActivatedSoftware[]>([])
const activationResult = ref<ActivationResult>({ status: 'success' })

// 生成时间戳
const getTimestamp = () => {
  const now = new Date()
  const year = now.getFullYear()
  const month = String(now.getMonth() + 1).padStart(2, '0')
  const day = String(now.getDate()).padStart(2, '0')
  const hours = String(now.getHours()).padStart(2, '0')
  const minutes = String(now.getMinutes()).padStart(2, '0')
  const seconds = String(now.getSeconds()).padStart(2, '0')
  return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`
}

// 复制激活码
const copyActivationCode = async () => {
  if (activationResult.value.activationCode) {
    try {
      await navigator.clipboard.writeText(activationResult.value.activationCode)
      alert('激活码已复制到剪贴板')
    } catch (err) {
      console.error('复制失败:', err)
      alert('复制失败，请手动复制')
    }
  }
}

// 开始激活
const startActivation = async () => {
  isActivating.value = true
  showResultModal.value = false

  // 模拟激活过程
  const steps = [
    '正在连接服务器...',
    '正在验证设备信息...',
    '正在应用激活...'
  ]

  for (const step of steps) {
    activationProgress.value = step
    await new Promise(resolve => setTimeout(resolve, 1000))
  }

  // 模拟激活结果（70% 成功率）
  const isSuccess = Math.random() > 0.3

  if (isSuccess) {
    // 模拟多个软件的激活
    const products = ['IntelliJIdea2025.2', 'WebStorm2025.2', 'PyCharm2025.2']
    const currentTime = getTimestamp()

    // 为每个产品添加到已激活软件列表
    products.forEach(product => {
      activatedSoftwareList.value.unshift({
        productName: product,
        time: currentTime,
        status: 'success'
      })
    })

    // 生成模拟激活码
    const activationCode = 'JETBRAINS-LICENSE-KEY-2025\n' +
      'ABCD1234-EFGH5678-IJKL9012-MNOP3456\n' +
      'QRST7890-UVWX1234-YZAB5678-CDEF9012\n' +
      'VALID-UNTIL: 2026-12-31'

    activationResult.value = {
      status: 'success',
      activationCode: activationCode
    }
  } else {
    activationResult.value = {
      status: 'failed',
      errorMessage: '激活失败，请重试或联系客服'
    }
  }

  isActivating.value = false
  showResultModal.value = true
}

// 重试激活
const retryActivation = () => {
  closeModal()
  setTimeout(() => {
    startActivation()
  }, 300)
}

// 关闭弹窗
const closeModal = () => {
  showResultModal.value = false
}

// 返回首页
const goBack = () => {
  window.history.back()
}
</script>