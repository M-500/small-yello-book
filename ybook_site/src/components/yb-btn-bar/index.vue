<template>
  <div class="container">
    <router-link class="item"
                 :class="active ? 'active' : ''"
                 @click="clickHdl"
                 :to="nextPath">
      <div class="icon">
        <img :src="imgUrl"
             :class="isImage ? 'avatar': ''"
             alt="" />
        <div class="count"
             v-if="countNumber > 0">{{ countNumber }}</div>
      </div>
      <div class="title">{{ title }}</div>
    </router-link>
  </div>
</template>

<script setup>
import { defineEmits } from 'vue';
const emit = defineEmits(['clickBtn']);
const clickHdl = (data) => {
  emit('clickBtn', data);
}
defineProps({
  barType: {
    type: String,
    default: 'link'
  },
  isImage: {
    type: Boolean,
    default: false
  },
  nextPath: {
    type: String,
    default: 'icon'
  },
  imgUrl: {
    type: String,
    default: ''
  },
  countNumber: {
    type: Number,
    default: 0
  },
  title: {
    type: String,
    default: ''
  },
  active: {
    type: Boolean,
    default: false
  }
})
</script>

<style lang="scss" scoped>
.container{
	.active{
		background-color: $active-color;
	}
	.item{
		display: flex;
		padding-left: 16px;
		margin-bottom: 8px;
		align-items: center;
		height: 45px;
		cursor: pointer;
		border-radius: 50px;
		text-decoration: none; // 去掉a标签的下划线
		.icon {
			position: relative;
			img {
				width: 24px;
				height: 24px;
			}
			.avatar{
				border-radius: 50%;
			}
			.count{
				position: absolute;
				top: 0;
				right: 0;
				background-color: $primary-color-red;
				color: white;
				display: flex;
				align-items: center;
				justify-content: center;
				padding: 0 4px;
				min-width: 16px;
				height: 16px;
				border-radius: 999px;
				font-size: 12px;
				font-weight: 500;
				transform: translate(calc(100% - 4px), calc(-100% + 4px));
				z-index: 1;
			}
		}
		.title{
			font-size: 16px;
			font-weight: 600;
			margin-left: 12px;
			color: $gray-color;
		}
		&:hover{
			background-color: $active-color;
		}
	}
}
</style>
