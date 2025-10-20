<script lang="ts" setup>
import {useDrvice} from "./store/useDrvice";

const drviceStore = useDrvice();
drviceStore.getDrviceInfo()
drviceStore.getCode()
</script>

<template>
  <router-view />
</template>

<style>
</style>
