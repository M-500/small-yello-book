<template>
  <div class="note-item">
    <div class="container">
      <a href=""></a>
      <router-link :to="'/note/detail/' + props.item.uuid"
                   class="cover">
        <img :src="props.item.cover"
             alt="">
      </router-link>
      <!-- <a href=""
         class="cover">
        <img :src="props.item.cover"
             alt="">
      </a> -->
      <note-footer :item="props.item" />
    </div>
  </div>
</template>

<script lang="ts" setup>
import NoteFooter from "@/components/note-footer/index.vue"
import { RouterLink } from "vue-router";
import { defineProps } from "vue";

const props:any = defineProps({
  item: {
    type: Object,
    required: true,
    default: () => ({}) as any,
  },
});
</script>

<style lang="scss" scoped>
.note-item{
	// position: absolute;
	// left: 0;
	// top: 0;
	// margin-right: 20px;
	width: 221.33px; // 这个需要后期需要根据屏幕来适配
	.container{
		width: 100%;
		height: 100%;
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		.cover{
			position: relative;
			width: 221.33px; // 这个需要后期需要根据屏幕来适配
			display: flex;
			border-radius: 16px;
			overflow: hidden;
			outline: 1px solid rgba(0,0,0,0.08);
			outline-offset: -1px;
			transition: background .2s;
			transform: translateZ(0);
		}
	}
}
@media (max-width: 768px){
	.note-item{
		width: 197px;
		.container{
			.cover{
				width: 100%;
			}
		}
	}
}
</style>
