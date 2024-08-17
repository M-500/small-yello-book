<template>
  <div class="">
    <router-link :to="detailURL">
      <img :src="imgURL"
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
});
</script>

<style lang="scss" scoped>
</style>
