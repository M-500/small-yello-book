<template>
  <div class="more-container">
    <ybBtnBar nextPath=""
              :imgUrl="moreIcon"
              @clickBtn="clickHandler"
              title="登出" />
  </div>
</template>

<script setup>
import moreIcon from '@/assets/icons/more.svg';
import ybBtnBar from "@/components/yb-btn-bar/index.vue";
import useUserStore from '@/stores/modules/user';
const useStore = useUserStore();

const clickHandler = () => {
  console.log('click more');
  useStore.userLogout();
}
</script>

<style lang="scss" scoped></style>
