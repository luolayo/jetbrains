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
                <span>只支持21年及以后的版本，之前的版本不支持，请关闭需要激活的软件</span>
              </li>
              <li class="flex items-start gap-2">
                <svg class="w-5 h-5 text-blue-600 flex-shrink-0 mt-0.5" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
                </svg>
                <span>激活成功后，需要手动输入激活码，激活码会在激活信息中告诉您，如果您已经打开了要激活的软件，请重新启动软件</span>
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

      <!-- 激活类型选择区域 -->
      <div class="bg-white rounded-2xl shadow-sm p-6 mb-6">
        <h2 class="text-lg font-semibold text-slate-900 mb-4">选择激活类型</h2>

        <!-- 提示信息 -->
        <div class="bg-amber-50 border-2 border-amber-200 rounded-xl p-4 mb-4">
          <div class="flex gap-3">
            <svg class="w-5 h-5 text-amber-600 flex-shrink-0 mt-0.5" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd" d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z" clip-rule="evenodd" />
            </svg>
            <div class="flex-1">
              <p class="text-sm font-medium text-amber-900 mb-1">提醒</p>
              <p class="text-sm text-amber-800">
                <strong>软件本体激活：</strong>激活 JetBrains 软件（IDEA、WebStorm 等）<br>
                <strong>插件激活：</strong>激活插件，输入插件名称即可获取激活码 <br>
                <strong>插件激活需要先使用本根据激活软件本体才能激活插件，否则无法激活，激活后不保证能否使用</strong>
              </p>
            </div>
          </div>
        </div>

        <!-- 激活类型选项 -->
        <div class="grid md:grid-cols-2 gap-4">
          <button
            @click="activationType = 'software'"
            :class="[
              'p-5 rounded-xl border-2 transition-all text-left',
              activationType === 'software'
                ? 'border-blue-500 bg-blue-50 shadow-md'
                : 'border-slate-200 hover:border-slate-300 hover:shadow-sm'
            ]"
          >
            <div class="flex items-start gap-3">
              <div :class="[
                'w-6 h-6 rounded-full border-2 flex items-center justify-center flex-shrink-0 mt-0.5',
                activationType === 'software' ? 'border-blue-500 bg-blue-500' : 'border-slate-300'
              ]">
                <div v-if="activationType === 'software'" class="w-2.5 h-2.5 bg-white rounded-full"></div>
              </div>
              <div class="flex-1">
                <h3 class="font-semibold text-slate-900 mb-1">软件本体激活</h3>
                <p class="text-sm text-slate-600">激活 JetBrains 全系软件</p>
                <div class="mt-2 flex items-center gap-1.5">
                  <span class="text-xs bg-green-100 text-green-700 px-2 py-0.5 rounded-full font-medium">
                    IDEA/WebStorm/PyCharm等
                  </span>
                </div>
              </div>
            </div>
          </button>

          <button
            @click="activationType = 'plugin'"
            :class="[
              'p-5 rounded-xl border-2 transition-all text-left',
              activationType === 'plugin'
                ? 'border-blue-500 bg-blue-50 shadow-md'
                : 'border-slate-200 hover:border-slate-300 hover:shadow-sm'
            ]"
          >
            <div class="flex items-start gap-3">
              <div :class="[
                'w-6 h-6 rounded-full border-2 flex items-center justify-center flex-shrink-0 mt-0.5',
                activationType === 'plugin' ? 'border-blue-500 bg-blue-500' : 'border-slate-300'
              ]">
                <div v-if="activationType === 'plugin'" class="w-2.5 h-2.5 bg-white rounded-full"></div>
              </div>
              <div class="flex-1">
                <h3 class="font-semibold text-slate-900 mb-1">插件激活</h3>
                <p class="text-sm text-slate-600">输入插件名称获取激活码</p>
                <div class="mt-2 flex items-center gap-1.5">
                  <span class="text-xs bg-purple-100 text-purple-700 px-2 py-0.5 rounded-full font-medium">
                    插件专用
                  </span>
                </div>
              </div>
            </div>
          </button>
        </div>

        <!-- 插件名称输入框 (仅在选择插件激活时显示) -->
        <div v-if="activationType === 'plugin'" class="mt-4">
          <label class="block text-sm font-semibold text-slate-900 mb-2">插件名称</label>
          <input
            v-model="pluginName"
            type="text"
            placeholder="请输入插件名称，例如：MyBatis Plugin"
            class="w-full px-4 py-3 border-2 border-slate-200 rounded-xl focus:outline-none focus:border-blue-500 transition-colors"
            @keyup.enter="searchPlugin"
          />
          <p class="mt-2 text-xs text-slate-500">按 Enter 键搜索插件</p>
        </div>
      </div>

      <!-- 激活按钮区域 -->
      <div class="bg-white rounded-2xl shadow-sm p-8 mb-6">
        <div class="text-center">
          <div class="mb-6">
            <div :class="[
              'w-20 h-20 rounded-2xl flex items-center justify-center mx-auto mb-4 shadow-lg',
              activationType === 'plugin'
                ? 'bg-gradient-to-br from-purple-500 to-purple-600'
                : 'bg-gradient-to-br from-blue-500 to-indigo-600'
            ]">
              <svg v-if="activationType === 'software'" class="w-10 h-10 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 7a2 2 0 012 2m4 0a6 6 0 01-7.743 5.743L11 17H9v2H7v2H4a1 1 0 01-1-1v-2.586a1 1 0 01.293-.707l5.964-5.964A6 6 0 1121 9z" />
              </svg>
              <svg v-else class="w-10 h-10 text-white" fill="currentColor" viewBox="0 0 20 20">
                <path d="M11 3a1 1 0 10-2 0v1a1 1 0 102 0V3zM15.657 5.757a1 1 0 00-1.414-1.414l-.707.707a1 1 0 001.414 1.414l.707-.707zM18 10a1 1 0 01-1 1h-1a1 1 0 110-2h1a1 1 0 011 1zM5.05 6.464A1 1 0 106.464 5.05l-.707-.707a1 1 0 00-1.414 1.414l.707.707zM5 10a1 1 0 01-1 1H3a1 1 0 110-2h1a1 1 0 011 1zM8 16v-1h4v1a2 2 0 11-4 0zM12 14c.015-.34.208-.646.477-.859a4 4 0 10-4.954 0c.27.213.462.519.477.859h4z" />
              </svg>
            </div>
            <h2 class="text-2xl font-bold text-slate-900 mb-2">
              {{ activationType === 'plugin' ? '准备获取插件激活码' : '准备激活' }}
            </h2>
            <p class="text-slate-600">
              {{ activationType === 'plugin' ? '点击下方按钮搜索插件并获取激活码' : '点击下方按钮开始激活您的 JetBrains 产品' }}
            </p>
          </div>

          <button
            @click="startActivation"
            :disabled="isActivating || (activationType === 'plugin' && !pluginName.trim())"
            :class="[
              'px-8 py-4 rounded-xl font-semibold text-lg transition-all shadow-lg',
              isActivating || (activationType === 'plugin' && !pluginName.trim())
                ? 'bg-slate-400 cursor-not-allowed'
                : activationType === 'plugin'
                  ? 'bg-gradient-to-r from-purple-600 to-purple-700 hover:from-purple-700 hover:to-purple-800 text-white hover:shadow-xl transform hover:scale-105'
                  : 'bg-gradient-to-r from-blue-600 to-indigo-600 hover:from-blue-700 hover:to-indigo-700 text-white hover:shadow-xl transform hover:scale-105'
            ]"
          >
            <span v-if="!isActivating">
              {{ activationType === 'plugin' ? '搜索插件' : '立即激活' }}
            </span>
            <span v-else class="flex items-center justify-center gap-3">
              <svg class="w-6 h-6 animate-spin" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
              {{ activationProgress }}
            </span>
          </button>

          <!-- 插件激活时的提示 -->
          <p v-if="activationType === 'plugin' && !pluginName.trim()" class="mt-3 text-sm text-slate-500">
            请先输入插件名称
          </p>
        </div>
      </div>

      <!-- 激活信息 -->
      <div v-if="activationType === 'software'" class="bg-white rounded-2xl shadow-sm p-6 mb-6">
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
                <h3 class="font-semibold text-slate-900 mb-1 truncate">{{ software.productName }} - {{software.prodcutVersion}}</h3>
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

      <!-- 已激活插件列表 -->
      <div v-if="activationType === 'plugin'" class="bg-white rounded-2xl shadow-sm p-6">
        <h2 class="text-lg font-semibold text-slate-900 mb-4">已激活插件</h2>
        <div v-if="activatedPluginList.length === 0" class="text-center py-8 text-slate-500">
          <svg class="w-16 h-16 mx-auto mb-3 text-slate-300" fill="currentColor" viewBox="0 0 20 20">
            <path d="M11 3a1 1 0 10-2 0v1a1 1 0 102 0V3zM15.657 5.757a1 1 0 00-1.414-1.414l-.707.707a1 1 0 001.414 1.414l.707-.707zM18 10a1 1 0 01-1 1h-1a1 1 0 110-2h1a1 1 0 011 1zM5.05 6.464A1 1 0 106.464 5.05l-.707-.707a1 1 0 00-1.414 1.414l.707.707zM5 10a1 1 0 01-1 1H3a1 1 0 110-2h1a1 1 0 011 1zM8 16v-1h4v1a2 2 0 11-4 0zM12 14c.015-.34.208-.646.477-.859a4 4 0 10-4.954 0c.27.213.462.519.477.859h4z" />
          </svg>
          <p>暂无已激活插件</p>
        </div>
        <div v-else class="grid md:grid-cols-2 gap-4">
          <!-- 每个插件一个独立卡片 -->
          <div
            v-for="(plugin, index) in activatedPluginList"
            :key="index"
            :class="[
              'p-4 border-2 rounded-xl transition-all',
              plugin.status === 'success'
                ? 'border-purple-200 bg-purple-50 hover:shadow-md'
                : 'border-red-200 bg-red-50'
            ]"
          >
            <div class="flex items-start gap-3">
              <div :class="[
                'w-10 h-10 rounded-lg flex items-center justify-center flex-shrink-0',
                plugin.status === 'success' ? 'bg-purple-100' : 'bg-red-100'
              ]">
                <svg v-if="plugin.status === 'success'" class="w-6 h-6 text-purple-600" fill="currentColor" viewBox="0 0 20 20">
                  <path d="M11 3a1 1 0 10-2 0v1a1 1 0 102 0V3zM15.657 5.757a1 1 0 00-1.414-1.414l-.707.707a1 1 0 001.414 1.414l.707-.707zM18 10a1 1 0 01-1 1h-1a1 1 0 110-2h1a1 1 0 011 1zM5.05 6.464A1 1 0 106.464 5.05l-.707-.707a1 1 0 00-1.414 1.414l.707.707zM5 10a1 1 0 01-1 1H3a1 1 0 110-2h1a1 1 0 011 1zM8 16v-1h4v1a2 2 0 11-4 0zM12 14c.015-.34.208-.646.477-.859a4 4 0 10-4.954 0c.27.213.462.519.477.859h4z" />
                </svg>
                <svg v-else class="w-6 h-6 text-red-600" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z" clip-rule="evenodd" />
                </svg>
              </div>
              <div class="flex-1 min-w-0">
                <h3 class="font-semibold text-slate-900 mb-1 truncate">{{ plugin.productName }}</h3>
                <div class="flex items-center gap-2 mb-1">
                  <span :class="[
                    'px-2 py-0.5 rounded-md text-xs font-semibold',
                    plugin.status === 'success' ? 'bg-purple-600 text-white' : 'bg-red-600 text-white'
                  ]">
                    {{ plugin.status === 'success' ? '激活成功' : '激活失败' }}
                  </span>
                  <span class="text-xs text-slate-500">{{ plugin.time }}</span>
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
              <div :class="[
                'w-20 h-20 rounded-full flex items-center justify-center mx-auto mb-4',
                activationType === 'plugin' ? 'bg-purple-100' : 'bg-green-100'
              ]">
                <svg v-if="activationType === 'software'" class="w-10 h-10 text-green-600" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
                </svg>
                <svg v-else class="w-10 h-10 text-purple-600" fill="currentColor" viewBox="0 0 20 20">
                  <path d="M11 3a1 1 0 10-2 0v1a1 1 0 102 0V3zM15.657 5.757a1 1 0 00-1.414-1.414l-.707.707a1 1 0 001.414 1.414l.707-.707zM18 10a1 1 0 01-1 1h-1a1 1 0 110-2h1a1 1 0 011 1zM5.05 6.464A1 1 0 106.464 5.05l-.707-.707a1 1 0 00-1.414 1.414l.707.707zM5 10a1 1 0 01-1 1H3a1 1 0 110-2h1a1 1 0 011 1zM8 16v-1h4v1a2 2 0 11-4 0zM12 14c.015-.34.208-.646.477-.859a4 4 0 10-4.954 0c.27.213.462.519.477.859h4z" />
                </svg>
              </div>
              <h3 class="text-2xl font-bold text-slate-900 mb-2">
                {{ activationType === 'plugin' ? '获取成功！' : '激活成功！' }}
              </h3>
              <p class="text-slate-600 mb-6">
                {{ activationType === 'plugin' ? '插件激活码已生成，请复制使用' : '请将激活码复制到软件里即可成功激活' }}
              </p>

              <!-- 激活码文本域 -->
              <div class="bg-slate-50 rounded-xl p-4 mb-6 text-left">
                <div class="flex items-center justify-between mb-2">
                  <h4 class="text-sm font-semibold text-slate-900">
                    {{ activationType === 'plugin' ? '插件激活码' : '激活码' }}
                  </h4>
                  <button
                    @click="copyActivationCode"
                    :class="[
                      'text-xs text-white px-3 py-1 rounded-md transition-colors',
                      activationType === 'plugin'
                        ? 'bg-purple-600 hover:bg-purple-700'
                        : 'bg-blue-600 hover:bg-blue-700'
                    ]"
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
                :class="[
                  'w-full py-3 text-white rounded-xl font-semibold transition-colors',
                  activationType === 'plugin'
                    ? 'bg-purple-600 hover:bg-purple-700'
                    : 'bg-green-600 hover:bg-green-700'
                ]"
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
              <h3 class="text-2xl font-bold text-slate-900 mb-2">
                {{ activationType === 'plugin' ? '获取失败' : '激活失败' }}
              </h3>
              <p class="text-slate-600 mb-6">{{ activationResult.errorMessage }}</p>

              <!-- 软件激活失败的原因提示 -->
              <div v-if="activationType === 'software'" class="bg-red-50 border-2 border-red-200 rounded-xl p-4 mb-6 text-left">
                <h4 class="font-semibold text-red-900 mb-2">可能的原因：</h4>
                <ul class="space-y-1 text-sm text-red-700">
                  <li>软件安装后开机后没有打开过，读取不到配置目录</li>
                  <li>win如果打开了软件，请把所有软件关掉再重启工具</li>
                  <li>mac 没有给磁盘权限</li>
                  <li>激活失败大概率是读取不到文件夹，没有文件夹权限，被杀毒拦截等，如果关闭杀毒或者按照上述操作还是不行，请去教程里的问题列表查找问题和解决办法</li>
                </ul>
              </div>

              <!-- 插件激活失败的原因提示 -->
              <div v-else class="bg-red-50 border-2 border-red-200 rounded-xl p-4 mb-6 text-left">
                <h4 class="font-semibold text-red-900 mb-2">可能的原因：</h4>
                <ul class="space-y-1 text-sm text-red-700">
                  <li>插件名称输入错误或不存在</li>
                  <li>该插件不需要激活</li>
                  <li>插件库中暂无此插件的激活信息</li>
                  <li>请检查插件名称是否正确</li>
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
                  :class="[
                    'flex-1 py-3 text-white rounded-xl font-semibold transition-colors',
                    activationType === 'plugin'
                      ? 'bg-purple-600 hover:bg-purple-700'
                      : 'bg-blue-600 hover:bg-blue-700'
                  ]"
                >
                  重试
                </button>
              </div>
            </div>
          </div>
        </transition>
      </div>
    </transition>

    <!-- 激活流程引导弹窗 -->
    <transition
      enter-active-class="transition ease-out duration-200"
      enter-from-class="opacity-0"
      enter-to-class="opacity-100"
      leave-active-class="transition ease-in duration-150"
      leave-from-class="opacity-100"
      leave-to-class="opacity-0"
    >
      <div
        v-if="showGuideModal"
        class="fixed inset-0 backdrop-blur-sm bg-white/30 flex items-center justify-center p-4 z-50"
      >
        <transition
          enter-active-class="transition ease-out duration-200"
          enter-from-class="opacity-0 scale-95"
          enter-to-class="opacity-100 scale-100"
          leave-active-class="transition ease-in duration-150"
          leave-from-class="opacity-100 scale-100"
          leave-to-class="opacity-0 scale-95"
        >
          <div
            v-if="showGuideModal"
            class="bg-white rounded-2xl shadow-xl max-w-lg w-full border border-slate-200"
            @click.stop
          >
            <!-- 头部 -->
            <div class="px-6 py-5 border-b border-slate-200">
              <div class="flex items-center justify-between">
                <div>
                  <h2 class="text-xl font-bold text-slate-900">激活流程说明</h2>
                  <p class="text-sm text-slate-500 mt-1">请按以下步骤操作以确保激活成功</p>
                </div>
                <button
                  @click="closeGuideModal"
                  class="w-8 h-8 rounded-lg hover:bg-slate-100 transition-colors flex items-center justify-center text-slate-400 hover:text-slate-600"
                >
                  <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                  </svg>
                </button>
              </div>
            </div>

            <!-- 内容区 -->
            <div class="px-6 py-5">
              <!-- 步骤列表 -->
              <div class="space-y-4">
                <!-- 步骤1 -->
                <div class="bg-gradient-to-r from-blue-50 to-indigo-50 rounded-xl p-4 border border-blue-100">
                  <div class="flex gap-3">
                    <div class="flex-shrink-0">
                      <div class="w-8 h-8 rounded-lg bg-gradient-to-br from-blue-500 to-indigo-600 text-white flex items-center justify-center font-bold">
                        1
                      </div>
                    </div>
                    <div class="flex-1">
                      <h3 class="font-semibold text-slate-900 mb-1">打开所有需要激活的软件并关闭</h3>
                      <p class="text-sm text-slate-600 leading-relaxed">
                        先启动一次您需要激活的 JetBrains 软件(如 IDEA、WebStorm 等),然后完全关闭它们。这一步是为了让工具能够读取软件的配置目录。
                      </p>
                    </div>
                  </div>
                </div>

                <!-- 步骤2 -->
                <div class="bg-gradient-to-r from-blue-50 to-indigo-50 rounded-xl p-4 border border-blue-100">
                  <div class="flex gap-3">
                    <div class="flex-shrink-0">
                      <div class="w-8 h-8 rounded-lg bg-gradient-to-br from-blue-500 to-indigo-600 text-white flex items-center justify-center font-bold">
                        2
                      </div>
                    </div>
                    <div class="flex-1">
                      <h3 class="font-semibold text-slate-900 mb-1">点击激活按钮</h3>
                      <p class="text-sm text-slate-600 leading-relaxed">
                        在本页面选择"软件本体激活",然后点击"立即激活"按钮。等待激活过程完成,系统会自动处理相关配置。提示成功后，请查看下面的激活列表里面有没有显示你需要激活的软件和版本。
                      </p>
                    </div>
                  </div>
                </div>

                <!-- 步骤3 -->
                <div class="bg-gradient-to-r from-blue-50 to-indigo-50 rounded-xl p-4 border border-blue-100">
                  <div class="flex gap-3">
                    <div class="flex-shrink-0">
                      <div class="w-8 h-8 rounded-lg bg-gradient-to-br from-blue-500 to-indigo-600 text-white flex items-center justify-center font-bold">
                        3
                      </div>
                    </div>
                    <div class="flex-1">
                      <h3 class="font-semibold text-slate-900 mb-1">重启需要激活的软件</h3>
                      <p class="text-sm text-slate-600 leading-relaxed">
                        激活完成后,重新启动您的 JetBrains 软件。首次启动时,软件会检测到新的激活配置
                      </p>
                    </div>
                  </div>
                </div>

                <!-- 步骤4 -->
                <div class="bg-gradient-to-r from-blue-50 to-indigo-50 rounded-xl p-4 border border-blue-100">
                  <div class="flex gap-3">
                    <div class="flex-shrink-0">
                      <div class="w-8 h-8 rounded-lg bg-gradient-to-br from-blue-500 to-indigo-600 text-white flex items-center justify-center font-bold">
                        4
                      </div>
                    </div>
                    <div class="flex-1">
                      <h3 class="font-semibold text-slate-900 mb-1">输入生成的激活码</h3>
                      <p class="text-sm text-slate-600 leading-relaxed">
                        在软件的激活窗口中,选择"Activation Code"方式,然后将系统生成的激活码粘贴进去,完成最终激活。
                      </p>
                    </div>
                  </div>
                </div>
              </div>

              <!-- 额外提示 -->
              <div class="mt-4 bg-amber-50 border border-amber-200 rounded-xl p-3">
                <div class="flex gap-2">
                  <svg class="w-4 h-4 text-amber-600 flex-shrink-0 mt-0.5" fill="currentColor" viewBox="0 0 20 20">
                    <path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z" clip-rule="evenodd" />
                  </svg>
                  <p class="text-xs text-amber-800">
                    <strong class="font-semibold">温馨提示：</strong>如果重启多次输入激活码无效,请尝试切换需要激活的软件的版本
                  </p>
                </div>
              </div>
            </div>

            <!-- 底部按钮 -->
            <div class="px-6 py-5 bg-slate-50 rounded-b-2xl border-t border-slate-200">
              <button
                @click="closeGuideModal"
                class="w-full py-3 bg-gradient-to-r from-blue-600 to-indigo-600 hover:from-blue-700 hover:to-indigo-700 text-white rounded-xl font-semibold transition-all shadow-sm hover:shadow"
              >
                我知道了,开始激活
              </button>
            </div>
          </div>
        </transition>
      </div>
    </transition>
  </div>
</template>

<script setup lang="ts">
import {ref} from 'vue'
import {Download, getConde, Verification} from "../api";
import {useDrvice} from "../store/useDrvice";
import {decrypt} from "../util/crypt";
import {Actions, CopyText, DownloadAndExtract} from '../../wailsjs/go/main/App'

interface ActivatedSoftware {
  productName: string
  prodcutVersion?: string
  time: string
  status: 'success' | 'failed'
}


interface ActivationResult {
  status: 'success' | 'failed'
  activationCode?: string
  errorMessage?: string
}

const drive = useDrvice()

// 状态
const activationType = ref<'software' | 'plugin'>('software') // 激活类型：软件或插件
const pluginName = ref('') // 插件名称
const selectedVersion = ref<'high' | 'low'>('high')
const isActivating = ref(false)
const activationProgress = ref('正在连接服务器...')
const showResultModal = ref(false)
const showGuideModal = ref(true) // 激活流程引导弹窗
const activatedSoftwareList = ref<ActivatedSoftware[]>([])
const activatedPluginList = ref<ActivatedSoftware[]>([]) // 已激活插件列表
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

// 搜索插件
const searchPlugin = async () => {
  if (!pluginName.value.trim()) {
    alert('请输入插件名称')
    return
  }
  // 如果选择的是插件激活，直接触发激活
  await startActivation()
}

// 开始激活（统一入口）
const startActivation = async () => {
  if (activationType.value === 'plugin') {
    await startPluginActivation()
  } else {
    await startSoftwareActivation()
  }
}

// 软件激活
const startSoftwareActivation = async () => {
  isActivating.value = true
  showResultModal.value = false

  const steps = [
    '正在连接服务器...',
    '正在验证设备信息...',
    '正在下载激活数据...',
    '正在应用激活...'
  ]

  try {
    activationProgress.value = steps[0]
    const result = await Verification(drive.drviceInfo.uuid, drive.drviceInfo.mac)
    const dataString = await decrypt(result.data, drive.code)
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

    // 软件激活逻辑
    activationProgress.value = steps[2]
    const version = selectedVersion.value === 'high' ? 'new' : 'old'
    const resultDown = await Download(drive.drviceInfo.uuid, drive.drviceInfo.mac, version)

    // 检查返回的是 JSON 还是文件
    if (resultDown instanceof Blob) {
      activationProgress.value = steps[3]

      // 将 Blob 转换为 ArrayBuffer 然后转为字节数组
      const arrayBuffer = await resultDown.arrayBuffer()
      const uint8Array = new Uint8Array(arrayBuffer)

      // 调用 Go 后端方法下载并解压到 WorkDir
      const filename = `jetbrains-activation-${version}.zip`
      await DownloadAndExtract(Array.from(uint8Array), filename)
      const ActionsData = await Actions()
      if (ActionsData?.error){
        activationResult.value = {
          status: 'failed',
          errorMessage: ActionsData.error.join('\n') || '激活失败，请稍后重试'
        }
      } else {
        // 激活成功
        const txt = await CopyText()
        activationResult.value = {
          status: 'success',
          activationCode: txt,
        }
        // 填充已激活软件列表
        activatedSoftwareList.value = ActionsData.product.map((item: { productName: string;
          productVersion: string}) => ({
          productName: item.productName,
          prodcutVersion: item.productVersion,
          time: getTimestamp(),
          status: 'success'
        }))
      }

      showResultModal.value = true
      isActivating.value = false
    } else {
      console.log(resultDown)
      activationResult.value = {
        status: 'failed',
        errorMessage: resultDown.message || '下载失败，请稍后重试'
      }
      showResultModal.value = true
      isActivating.value = false
    }
  } catch (error) {
    console.error('软件激活失败:', error)
    activationResult.value = {
      status: 'failed',
      errorMessage: error instanceof Error ? error.message : '请把所有jetbrains关闭后重启工具再重试'
    }
    showResultModal.value = true
    isActivating.value = false
  }
}

// 插件激活
const startPluginActivation = async () => {
  // 检查插件名称
  if (!pluginName.value.trim()) {
    alert('请输入插件名称')
    return
  }

  isActivating.value = true
  showResultModal.value = false

  const steps = [
    '正在连接服务器...',
    '正在搜索插件...',
    '正在获取激活码...'
  ]

  try {
    activationProgress.value = steps[0]
    activationProgress.value = steps[1]
    const res = await getConde(pluginName.value.trim())
    if (res.code !== 200) {
      activationResult.value = {
        status: 'failed',
        errorMessage:  res.message || '获取插件激活码失败，请稍后重试'
      }
      showResultModal.value = true
      isActivating.value = false
      return
    }
    activationProgress.value = steps[2]

    // 激活成功
    activationResult.value = {
      status: 'success',
      activationCode: res.data
    }

    // 添加到已激活插件列表
    activatedPluginList.value.push({
      productName: pluginName.value.trim(),
      prodcutVersion: '',
      time: getTimestamp(),
      status: 'success'
    })

    showResultModal.value = true
    isActivating.value = false
  } catch (error) {
    activationResult.value = {
      status: 'failed',
      errorMessage: error instanceof Error ? error.message : '获取插件激活码失败，请稍后重试'
    }
    showResultModal.value = true
    isActivating.value = false
  }
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

// 关闭引导弹窗
const closeGuideModal = () => {
  showGuideModal.value = false
}

// 返回首页
const goBack = () => {
  window.history.back()
}
</script>