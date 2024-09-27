<template>
  <div class="more-container">
    <ybBtnBar nextPath="/more"
              :imgUrl="moreIcon"
              title="更多" />
  </div>
</template>

<script setup>
import moreIcon from '@/assets/icons/more.svg';
import ybBtnBar from "@/components/yb-btn-bar/index.vue";
</script>

<style lang="scss" scoped></style>
