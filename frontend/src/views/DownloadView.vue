<template>
  <div class="min-h-screen bg-gradient-to-br from-slate-50 to-slate-100">
    <div class="container mx-auto px-4 py-8 max-w-6xl">
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
        <h1 class="text-4xl font-bold text-slate-900 mb-2">下载 JetBrains 软件</h1>
        <p class="text-slate-600">选择您需要的产品、版本和操作系统</p>
      </div>

      <!-- 产品选择 -->
      <div class="bg-white rounded-2xl shadow-sm p-6 mb-6">
        <h2 class="text-lg font-semibold text-slate-900 mb-4">选择产品</h2>
        <div class="grid grid-cols-2 md:grid-cols-4 gap-3">
          <button
            v-for="product in products"
            :key="product.id"
            @click="selectedProduct = product.id; loadVersions()"
            :class="[
              'p-4 rounded-xl border-2 transition-all duration-200',
              selectedProduct === product.id
                ? 'border-blue-500 bg-blue-50 shadow-md'
                : 'border-slate-200 hover:border-slate-300 hover:shadow-sm'
            ]"
          >
            <div class="text-2xl mb-2">{{ product.icon }}</div>
            <div class="font-medium text-slate-900 text-sm">{{ product.name }}</div>
          </button>
        </div>
      </div>

      <!-- 版本和系统选择 -->
      <div class="bg-white rounded-2xl shadow-sm p-6 mb-6">
        <div class="grid md:grid-cols-2 gap-6">
          <!-- 版本选择 -->
          <div>
            <h2 class="text-lg font-semibold text-slate-900 mb-4">选择版本</h2>

            <!-- 版本类型切换 -->
            <div class="flex gap-2 mb-4">
              <button
                @click="versionType = 'latest'"
                :class="[
                  'flex-1 py-2 px-4 rounded-lg font-medium transition-all',
                  versionType === 'latest'
                    ? 'bg-blue-500 text-white shadow-md'
                    : 'bg-slate-100 text-slate-600 hover:bg-slate-200'
                ]"
              >
                最新版本
              </button>
              <button
                @click="versionType = 'history'; loadVersions()"
                :class="[
                  'flex-1 py-2 px-4 rounded-lg font-medium transition-all',
                  versionType === 'history'
                    ? 'bg-blue-500 text-white shadow-md'
                    : 'bg-slate-100 text-slate-600 hover:bg-slate-200'
                ]"
              >
                历史版本
              </button>
            </div>

            <!-- 历史版本选择器 -->
            <div v-if="versionType === 'history'" class="relative">
              <div v-if="loadingVersions" class="flex items-center justify-center py-8">
                <div class="flex flex-col items-center gap-3">
                  <div class="w-10 h-10 border-4 border-blue-500 border-t-transparent rounded-full animate-spin"></div>
                  <p class="text-sm text-slate-600">加载版本列表中...</p>
                </div>
              </div>

              <div v-else>
                <button
                  @click="showVersionDropdown = !showVersionDropdown"
                  class="w-full flex items-center justify-between px-4 py-3 bg-slate-50 border-2 border-slate-200 rounded-xl hover:border-slate-300 transition-colors"
                >
                  <span class="font-medium text-slate-900">{{ selectedVersion }}</span>
                  <svg
                    :class="['w-5 h-5 text-slate-600 transition-transform', showVersionDropdown && 'rotate-180']"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                  >
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
                  </svg>
                </button>

                <!-- 下拉列表 -->
                <transition
                  enter-active-class="transition ease-out duration-200"
                  enter-from-class="opacity-0 translate-y-1"
                  enter-to-class="opacity-100 translate-y-0"
                  leave-active-class="transition ease-in duration-150"
                  leave-from-class="opacity-100 translate-y-0"
                  leave-to-class="opacity-0 translate-y-1"
                >
                  <div
                    v-if="showVersionDropdown"
                    class="absolute z-10 w-full mt-2 bg-white border-2 border-slate-200 rounded-xl shadow-lg max-h-64 overflow-y-auto"
                  >
                    <button
                      v-for="version in availableVersions"
                      :key="version"
                      @click="selectVersion(version)"
                      :class="[
                        'w-full px-4 py-3 text-left hover:bg-slate-50 transition-colors flex items-center justify-between',
                        selectedVersion === version && 'bg-blue-50 text-blue-600 font-medium'
                      ]"
                    >
                      <span>{{ version }}</span>
                      <svg
                        v-if="selectedVersion === version"
                        class="w-5 h-5"
                        fill="currentColor"
                        viewBox="0 0 20 20"
                      >
                        <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd" />
                      </svg>
                    </button>
                  </div>
                </transition>
              </div>
            </div>

            <!-- 最新版本显示 -->
            <div v-else class="px-4 py-3 bg-gradient-to-r from-blue-50 to-indigo-50 border-2 border-blue-200 rounded-xl">
              <div class="flex items-center gap-2">
                <svg class="w-5 h-5 text-blue-600" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
                </svg>
                <span class="font-semibold text-blue-900">{{ selectedVersion }}</span>
                <span class="text-xs bg-blue-600 text-white px-2 py-0.5 rounded-full">最新</span>
              </div>
            </div>
          </div>

          <!-- 系统选择 -->
          <div>
            <h2 class="text-lg font-semibold text-slate-900 mb-4">选择操作系统</h2>
            <div class="space-y-2">
              <button
                v-for="os in operatingSystems"
                :key="os.id"
                @click="selectedOS = os.id"
                :class="[
                  'w-full flex items-center gap-3 p-3 rounded-xl border-2 transition-all',
                  selectedOS === os.id
                    ? 'border-blue-500 bg-blue-50 shadow-md'
                    : 'border-slate-200 hover:border-slate-300 hover:shadow-sm'
                ]"
              >
                <span class="text-2xl">{{ os.icon }}</span>
                <span class="font-medium text-slate-900">{{ os.name }}</span>
                <span v-if="os.id === detectedOS" class="ml-auto text-xs bg-green-100 text-green-700 px-2 py-1 rounded-full">
                  已检测
                </span>
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- 下载选项 -->
      <div class="bg-white rounded-2xl shadow-sm p-6">
        <h2 class="text-lg font-semibold text-slate-900 mb-4">可用下载</h2>
        <div class="space-y-3">
          <div
            v-for="option in currentDownloadOptions"
            :key="option.link"
            class="flex items-center justify-between p-4 border-2 border-slate-200 rounded-xl hover:border-blue-300 hover:shadow-md transition-all group"
          >
            <div class="flex-1">
              <div class="flex items-center gap-2 mb-1">
                <h3 class="font-semibold text-slate-900">{{ option.name }}</h3>
                <span v-if="option.recommended" class="text-xs bg-blue-600 text-white px-2 py-0.5 rounded-full">
                  推荐
                </span>
              </div>
              <div class="flex items-center gap-4 text-sm text-slate-600">
                <span>{{ option.size }}</span>
                <span>{{ option.format }}</span>
              </div>
            </div>
            <button
              @click="downloadFile(option.link)"
              class="px-6 py-2.5 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors font-medium shadow-sm hover:shadow-md"
            >
              下载
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'

// 产品列表
const products = [
  { id: 'idea', name: 'IntelliJ IDEA', icon: '💡' },
  { id: 'pycharm', name: 'PyCharm', icon: '🐍' },
  { id: 'webstorm', name: 'WebStorm', icon: '🌐' },
  { id: 'phpstorm', name: 'PhpStorm', icon: '🐘' },
  { id: 'goland', name: 'GoLand', icon: '🔷' },
  { id: 'rider', name: 'Rider', icon: '🎮' },
  { id: 'clion', name: 'CLion', icon: '⚙️' },
  { id: 'datagrip', name: 'DataGrip', icon: '🗄️' }
]

// 操作系统列表
const operatingSystems = [
  { id: 'windows', name: 'Windows', icon: '🪟' },
  { id: 'mac', name: 'macOS', icon: '🍎' },
  { id: 'linux', name: 'Linux', icon: '🐧' }
]

// 状态
const selectedProduct = ref('idea')
const versionType = ref<'latest' | 'history'>('latest')
const selectedVersion = ref('2025.2.3')
const selectedOS = ref('windows')
const detectedOS = ref('windows')
const showVersionDropdown = ref(false)
const loadingVersions = ref(false)
const availableVersions = ref<string[]>([])

// 模拟网络请求加载版本列表
const loadVersions = async () => {
  if (versionType.value === 'history') {
    loadingVersions.value = true
    showVersionDropdown.value = false

    // 模拟网络延迟
    await new Promise(resolve => setTimeout(resolve, 1000))

    // 模拟从服务器获取的版本列表
    availableVersions.value = [
      '2025.2.3',
      '2025.2.2',
      '2025.2.1',
      '2025.2.0',
      '2025.1.5',
      '2025.1.4',
      '2025.1.3',
      '2024.3.2',
      '2024.3.1',
      '2024.3.0'
    ]

    loadingVersions.value = false
  }
}

// 选择版本
const selectVersion = (version: string) => {
  selectedVersion.value = version
  showVersionDropdown.value = false
}

// 检测操作系统
onMounted(() => {
  const userAgent = navigator.userAgent.toLowerCase()
  if (userAgent.includes('mac')) {
    detectedOS.value = 'mac'
    selectedOS.value = 'mac'
  } else if (userAgent.includes('linux')) {
    detectedOS.value = 'linux'
    selectedOS.value = 'linux'
  } else {
    detectedOS.value = 'windows'
    selectedOS.value = 'windows'
  }
})

// 下载选项数据（基于您提供的 IntelliJ IDEA 数据）
const downloadData = {
  windows: [
    { name: 'Windows (x64)', size: '1.1 GB', format: '.exe', link: 'https://download.jetbrains.com/idea/ideaIU-2025.2.3.exe', recommended: true },
    { name: 'Windows (ARM64)', size: '1.0 GB', format: '.exe', link: 'https://download.jetbrains.com/idea/ideaIU-2025.2.3-aarch64.exe', recommended: false }
  ],
  mac: [
    { name: 'macOS (Apple Silicon)', size: '1.2 GB', format: '.dmg', link: 'https://download.jetbrains.com/idea/ideaIU-2025.2.3-aarch64.dmg', recommended: true },
    { name: 'macOS (Intel)', size: '1.2 GB', format: '.dmg', link: 'https://download.jetbrains.com/idea/ideaIU-2025.2.3.dmg', recommended: false }
  ],
  linux: [
    { name: 'Linux (x64)', size: '1.1 GB', format: '.tar.gz', link: 'https://download.jetbrains.com/idea/ideaIU-2025.2.3.tar.gz', recommended: true },
    { name: 'Linux (ARM64)', size: '1.0 GB', format: '.tar.gz', link: 'https://download.jetbrains.com/idea/ideaIU-2025.2.3-aarch64.tar.gz', recommended: false }
  ]
}

// 当前下载选项
const currentDownloadOptions = computed(() => {
  return downloadData[selectedOS.value as keyof typeof downloadData] || []
})

// 下载文件
const downloadFile = (url: string) => {
  window.open(url, '_blank')
}

// 返回首页
const goBack = () => {
  window.history.back()
}
</script>