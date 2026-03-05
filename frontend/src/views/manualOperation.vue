<template>
  <div class="min-h-screen bg-gradient-to-br from-slate-50 to-slate-100">
    <div class="container mx-auto px-4 py-8 max-w-4xl">
      <!-- 返回按钮 -->
      <button
        @click="goBack"
        class="mb-6 flex items-center gap-2 text-slate-600 hover:text-slate-900 transition-colors"
      >
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"/>
        </svg>
        返回激活页
      </button>

      <!-- 页面标题 -->
      <div class="text-center mb-8">
        <h1 class="text-4xl font-bold text-slate-900 mb-2">手动识别软件</h1>
        <p class="text-slate-600">适用于自动识别失败的 Windows 用户</p>
      </div>

      <!-- 提示卡片 -->
      <div class="bg-gradient-to-r from-amber-50 to-orange-50 border-2 border-amber-200 rounded-2xl p-6 mb-6">
        <div class="flex gap-4">
          <div class="flex-shrink-0">
            <div class="w-12 h-12 bg-amber-500 rounded-xl flex items-center justify-center">
              <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                      d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
              </svg>
            </div>
          </div>
          <div class="flex-1">
            <h3 class="text-lg font-semibold text-slate-900 mb-2">使用说明</h3>
            <ul class="space-y-2 text-sm text-slate-700">
              <li class="flex items-start gap-2">
                <svg class="w-5 h-5 text-amber-600 flex-shrink-0 mt-0.5" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd"
                        d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z"
                        clip-rule="evenodd"/>
                </svg>
                <span>此功能仅适用于 <strong>Windows</strong> 系统，当自动识别无法找到已安装的软件时使用</span>
              </li>
              <li class="flex items-start gap-2">
                <svg class="w-5 h-5 text-amber-600 flex-shrink-0 mt-0.5" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd"
                        d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z"
                        clip-rule="evenodd"/>
                </svg>
                <span>请选择 JetBrains 软件的 <strong>安装根目录</strong>（例如：C:\Program Files\JetBrains\IntelliJ IDEA 2024.1）</span>
              </li>
              <li class="flex items-start gap-2">
                <svg class="w-5 h-5 text-amber-600 flex-shrink-0 mt-0.5" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd"
                        d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z"
                        clip-rule="evenodd"/>
                </svg>
                <span>激活前请确保已关闭需要激活的 JetBrains 软件</span>
              </li>
              <li class="flex items-start gap-2 bg-amber-100/60 rounded-lg px-3 py-2 -mx-1">
                <svg class="w-5 h-5 text-amber-500 flex-shrink-0 mt-0.5" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd"
                        d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 10-2 0 1 1 0 002 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z"
                        clip-rule="evenodd"/>
                </svg>
                <span class="text-amber-800">该功能为新增功能，暂时不做保证，识别不到一定是自己改了配置文件目录或者不是安装版</span>
              </li>
            </ul>
          </div>
        </div>
      </div>

      <!-- 选择文件夹区域 -->
      <div class="bg-white rounded-2xl shadow-sm p-6 mb-6">
        <h2 class="text-lg font-semibold text-slate-900 mb-4">选择软件安装目录</h2>

        <div
          @click="selectFolder"
          class="border-2 border-dashed border-slate-300 rounded-xl p-8 text-center cursor-pointer hover:border-blue-400 hover:bg-blue-50 transition-all"
        >
          <div v-if="!selectedPath">
            <svg class="w-16 h-16 mx-auto mb-3 text-slate-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                    d="M3 7v10a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-6l-2-2H5a2 2 0 00-2 2z"/>
            </svg>
            <p class="text-slate-500 text-lg font-medium mb-1">点击选择软件安装目录</p>
            <p class="text-slate-400 text-sm">请选择 JetBrains 产品的安装根目录</p>
          </div>
          <div v-else class="flex items-center justify-center gap-3">
            <svg class="w-8 h-8 text-blue-500 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                    d="M3 7v10a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-6l-2-2H5a2 2 0 00-2 2z"/>
            </svg>
            <div class="text-left">
              <p class="text-slate-900 font-medium">已选择目录</p>
              <p class="text-blue-600 text-sm break-all">{{ selectedPath }}</p>
            </div>
            <svg class="w-5 h-5 text-slate-400 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                    d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z"/>
            </svg>
          </div>
        </div>
      </div>

      <!-- 激活按钮区域 -->
      <div class="bg-white rounded-2xl shadow-sm p-8 mb-6">
        <div class="text-center">
          <div class="mb-6">
            <div class="w-20 h-20 rounded-2xl bg-gradient-to-br from-green-500 to-emerald-600 flex items-center justify-center mx-auto mb-4 shadow-lg">
              <svg class="w-10 h-10 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                      d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z"/>
              </svg>
            </div>
            <h2 class="text-2xl font-bold text-slate-900 mb-2">手动激活</h2>
            <p class="text-slate-600">选择安装目录后点击下方按钮开始激活</p>
          </div>

          <button
            @click="startManualActivation"
            :disabled="!selectedPath || isActivating"
            :class="[
              'px-8 py-4 rounded-xl font-semibold text-lg transition-all shadow-lg',
              !selectedPath || isActivating
                ? 'bg-slate-400 cursor-not-allowed text-white'
                : 'bg-gradient-to-r from-green-600 to-emerald-600 hover:from-green-700 hover:to-emerald-700 text-white hover:shadow-xl transform hover:scale-105'
            ]"
          >
            <span v-if="!isActivating">开始激活</span>
            <span v-else class="flex items-center justify-center gap-3">
              <svg class="w-6 h-6 animate-spin" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor"
                      d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
              {{ activationProgress }}
            </span>
          </button>

          <p v-if="!selectedPath" class="mt-3 text-sm text-slate-500">
            请先选择软件安装目录
          </p>
        </div>
      </div>

      <!-- 已激活软件列表 -->
      <div class="bg-white rounded-2xl shadow-sm p-6 mb-6">
        <h2 class="text-lg font-semibold text-slate-900 mb-4">已激活软件</h2>

        <div v-if="activatedSoftwareList.length === 0" class="text-center py-8 text-slate-500">
          <svg class="w-16 h-16 mx-auto mb-3 text-slate-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                  d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"/>
          </svg>
          <p>暂无已激活软件</p>
        </div>
        <div v-else class="grid md:grid-cols-2 gap-4">
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
                <svg v-if="software.status === 'success'" class="w-6 h-6 text-green-600" fill="currentColor"
                     viewBox="0 0 20 20">
                  <path fill-rule="evenodd"
                        d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z"
                        clip-rule="evenodd"/>
                </svg>
                <svg v-else class="w-6 h-6 text-red-600" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd"
                        d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z"
                        clip-rule="evenodd"/>
                </svg>
              </div>
              <div class="flex-1 min-w-0">
                <h3 class="font-semibold text-slate-900 mb-1 truncate">{{ software.productName }}
                  <span v-if="software.productVersion"> - {{ software.productVersion }}</span>
                </h3>
                <div class="flex items-center gap-2 mb-1">
                  <span :class="[
                    'px-2 py-0.5 rounded-md text-xs font-semibold',
                    software.status === 'success' ? 'bg-green-600 text-white' : 'bg-red-600 text-white'
                  ]">
                    {{ software.status === 'success' ? '激活成功' : '激活失败' }}
                  </span>
                  <span class="text-xs text-slate-500">{{ software.time }}</span>
                </div>
                <!-- 复制激活码按钮 -->
                <button
                  v-if="software.status === 'success'"
                  @click="showActivationCodeModal"
                  class="mt-2 text-sm text-blue-600 hover:text-blue-800 hover:underline transition-colors"
                >
                  查看并复制激活码
                </button>
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
                  <path fill-rule="evenodd"
                        d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z"
                        clip-rule="evenodd"/>
                </svg>
              </div>
              <h3 class="text-2xl font-bold text-slate-900 mb-2">激活成功！</h3>
              <p class="text-slate-600 mb-6">
                请复制下方激活码，在软件激活窗口中选择 "Activation Code" 方式粘贴使用
              </p>

              <!-- 激活码文本域 -->
              <div class="bg-slate-50 rounded-xl p-4 mb-6 text-left">
                <div class="flex items-center justify-between mb-2">
                  <h4 class="text-sm font-semibold text-slate-900">激活码</h4>
                  <button
                    @click="copyActivationCode"
                    class="text-xs text-white px-3 py-1 rounded-md bg-green-600 hover:bg-green-700 transition-colors"
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
                class="w-full py-3 bg-green-600 hover:bg-green-700 text-white rounded-xl font-semibold transition-colors"
              >
                完成
              </button>
            </div>

            <!-- 失败状态 -->
            <div v-else class="text-center">
              <div class="w-20 h-20 bg-red-100 rounded-full flex items-center justify-center mx-auto mb-4">
                <svg class="w-10 h-10 text-red-600" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd"
                        d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z"
                        clip-rule="evenodd"/>
                </svg>
              </div>
              <h3 class="text-2xl font-bold text-slate-900 mb-2">激活失败</h3>
              <p class="text-slate-600 mb-6">{{ activationResult.errorMessage }}</p>

              <div class="bg-red-50 border-2 border-red-200 rounded-xl p-4 mb-6 text-left">
                <h4 class="font-semibold text-red-900 mb-2">可能的原因：</h4>
                <ul class="space-y-1 text-sm text-red-700">
                  <li>选择的目录不是 JetBrains 软件的安装目录</li>
                  <li>软件安装后没有打开过，缺少必要的配置文件</li>
                  <li>需要激活的软件没有完全关闭</li>
                  <li>被杀毒软件拦截，请尝试关闭杀毒软件后重试</li>
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
                  class="flex-1 py-3 bg-green-600 hover:bg-green-700 text-white rounded-xl font-semibold transition-colors"
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
import {ref} from 'vue'
import {useRouter} from 'vue-router'
import {Download, Verification} from "../api";
import {useDrvice} from "../store/useDrvice";
import {decrypt} from "../util/crypt";
import {CopyText, DownloadAndExtract, ManualActions, SelectDirectory} from '../../wailsjs/go/main/App'

interface ActivatedSoftware {
  productName: string
  productVersion?: string
  time: string
  status: 'success' | 'failed'
}

interface ActivationResult {
  status: 'success' | 'failed'
  activationCode?: string
  errorMessage?: string
}

const router = useRouter()
const drive = useDrvice()

const selectedPath = ref('')
const isActivating = ref(false)
const activationProgress = ref('正在处理...')
const showResultModal = ref(false)
const activatedSoftwareList = ref<ActivatedSoftware[]>([])
const activationResult = ref<ActivationResult>({status: 'success'})

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

// 选择文件夹
const selectFolder = async () => {
  try {
    const path = await SelectDirectory()
    if (path) {
      selectedPath.value = path
    }
  } catch (err) {
    console.error('选择目录失败:', err)
  }
}

// 开始手动激活
const startManualActivation = async () => {
  if (!selectedPath.value) return

  isActivating.value = true
  showResultModal.value = false

  const steps = [
    '正在验证安装目录...',
    '正在连接服务器...',
    '正在验证设备信息...',
    '正在下载激活数据...',
    '正在应用激活...'
  ]

  try {
    // 第一次调用：验证目录并检查是否需要下载
    activationProgress.value = steps[0]
    let result = await ManualActions(selectedPath.value)

    if (result.error) {
      activationResult.value = {
        status: 'failed',
        errorMessage: result.error
      }
      isActivating.value = false
      showResultModal.value = true
      return
    }

    if (result.needDownload) {
      // 需要下载：走验证 + 下载流程
      activationProgress.value = steps[1]
      const verifyResult = await Verification(drive.drviceInfo.uuid, drive.drviceInfo.mac)
      const dataString = await decrypt(verifyResult.data, drive.code)
      const data = JSON.parse(dataString)

      if (data.code !== 200) {
        activationResult.value = {
          status: 'failed',
          errorMessage: data.message || '设备校验失败，请稍后重试'
        }
        isActivating.value = false
        showResultModal.value = true
        return
      }

      activationProgress.value = steps[3]
      const resultDown = await Download(drive.drviceInfo.uuid, drive.drviceInfo.mac, 'new')

      if (resultDown instanceof Blob) {
        const arrayBuffer = await resultDown.arrayBuffer()
        const uint8Array = new Uint8Array(arrayBuffer)
        const filename = `jetbrains-activation-manual.zip`
        await DownloadAndExtract(Array.from(uint8Array), filename)

        // 下载完成后再次调用 ManualActions 执行激活
        activationProgress.value = steps[4]
        result = await ManualActions(selectedPath.value)

        if (result.error) {
          activationResult.value = {
            status: 'failed',
            errorMessage: result.error
          }
          isActivating.value = false
          showResultModal.value = true
          return
        }
      } else {
        activationResult.value = {
          status: 'failed',
          errorMessage: resultDown.message || '下载失败，请稍后重试'
        }
        showResultModal.value = true
        isActivating.value = false
        return
      }
    }

    // 激活成功，获取激活码
    const txt = await CopyText()
    activationResult.value = {
      status: 'success',
      activationCode: txt,
    }

    activatedSoftwareList.value.push({
      productName: result.product.productName || '未知产品',
      productVersion: result.product.productVersion || '',
      time: getTimestamp(),
      status: 'success'
    })

    showResultModal.value = true
    isActivating.value = false
  } catch (error) {
    console.error('手动激活失败:', error)
    activationResult.value = {
      status: 'failed',
      errorMessage: error instanceof Error ? error.message : '激活失败，请检查目录是否正确后重试'
    }
    showResultModal.value = true
    isActivating.value = false
  }
}

// 显示激活码弹窗
const showActivationCodeModal = () => {
  showResultModal.value = true
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

// 重试
const retryActivation = () => {
  closeModal()
  setTimeout(() => {
    startManualActivation()
  }, 300)
}

// 关闭弹窗
const closeModal = () => {
  showResultModal.value = false
}

// 返回
const goBack = () => {
  router.push('/activation')
}
</script>