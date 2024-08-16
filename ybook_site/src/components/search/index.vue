<template>
  <div class="search_box">
    <input type="text"
           v-model="searchValue"
           class="search_input"
           placeholder="搜索小红书">
    <div class="input_button">
      <div class="close_icon"
           @click="clearSearch"
           v-show="showClose">
        <svg-icon name="close"
                  color="#000" />
      </div>
      <div class="search_icon"
           @click="search">
        <svg-icon name="search"
                  color="#000" />
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, watch } from "vue"
// 搜索小红书  登陆探索更多内容
const searchValue = ref("")
const showClose = ref(false)

// 监听searchValue的变化,一旦有数据就将showClose设置为true
watch(searchValue, (val) => {
  showClose.value = (val.length > 0)
})

function clearSearch () {
  searchValue.value = ""
}

function search () {
  console.log("search")
}
</script>

<style lang="scss" scoped>
.search_box{
	width: calc(-33.6px + 40vw);
	height: 40px;
	position: fixed;
	left: 50%;
	transform: translate(-50%);
	.search_input{
		padding: 0 84px 0 16px;
    width: 100%;
    height: 100%;
    font-size: 16px;
    line-height: 120%;
		color: $gray-color;
		caret-color:$primary-color-red;
		background: rgba(0, 0, 0, 0.03);
		border-radius: 100px;
		// 不要边框
		border: none;
	}
	.input_button{
		// 定位到父级菜单的右边
		position: absolute;
		right: 0;
		top: 30%;
		display: flex;
		justify-content: flex-start;
		.search_icon{
			width: 40px;
			height: 100%;
			display: flex;
			align-items: center;
			justify-content: center;
			cursor: pointer;
		}
		.close_icon{
			width: 40px;
			height: 100%;
			display: flex;
			align-items: center;
			justify-content: center;
			cursor: pointer;
		}
	}
}
</style>
