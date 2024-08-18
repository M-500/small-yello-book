<template>
  <div class="avatar_container">
    <router-link :to="detailURL"
                 class="user_avatar">
      <img :src="imgURL"
           :style="{width,height}"
           alt="">
    </router-link>

  </div>
</template>

<script setup>
import { RouterLink } from 'vue-router';
import { defineProps } from 'vue';

defineProps({
  avatar: String,
  detailURL: String,
  imgURL: String,
  name: String,
  id: String,
  ip: String,
  desc: String,
  age: String,
  follow: String,
  fans: String,
  like: String,
  width: {
    type: String,
    default: '48px'
  },
  height: {
    type: String,
    default: '48px'
  }
});
</script>

<style lang="scss" scoped>
.avatar_container{
  .user_avatar{
    margin-right: 24px;
    flex-shrink: 0;
    height: 100%;
    text-decoration: none;
    background-color: transparent;
    img{
      display: flex;
      align-items: center;
      justify-content: center;
      cursor: pointer;
      border-radius: 100%;
      border: 1px solid rgba(0,0,0,0.08);
      object-fit: cover;
    }
  }
}
</style>
