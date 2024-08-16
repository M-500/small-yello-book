<template>
  <div class="search_box">
    <input type="text"
           v-model="searchValue"
           class="search_input"
           placeholder="搜索小黄书">
    <div class="input_button">
      <div class="close_icon"
           @click="clearSearch"
           v-show="showClose">
        <svg-icon name="close"
                  width="16"
                  height="16"
                  color="#000" />
      </div>
      <div class="search_icon"
           @click="search">
        <svg-icon name="search"
                  width="16"
                  height="16"
                  color="#000" />
      </div>
    </div>

    <!-- 搜索关键词展示区域 -->
    <div class="sug-container-wrapper sug_pad"
         v-show="showClose">
      <div class="sug-container">
        <div class="sug-box">
          <div class="sug-wrapper">
            <div class="sug-item active">

              <div>
                <div><span class="highlight">心</span><span class="highlight">形</span><span class="">折</span><span class="">纸</span><span class="">教</span><span class="">程</span></div>
              </div>
            </div>
            <div class="sug-item">

              <div>
                <div><span class="highlight">心</span><span class="highlight">形</span><span class="">球</span><span class="">兰</span></div>
              </div>
            </div>
            <div class="sug-item">

              <div>
                <div><span class="highlight">心</span><span class="highlight">形</span><span class="">蛋</span><span class="">糕</span></div>
              </div>
            </div>
            <div class="sug-item">

              <div>
                <div><span class="highlight">心</span><span class="highlight">形</span><span class="">湖</span></div>
              </div>
            </div>
            <div class="sug-item">

              <div>
                <div><span class="highlight">心</span><span class="highlight">形</span><span class="">函</span><span class="">数</span></div>
              </div>
            </div>
            <div class="sug-item">

              <div>
                <div><span class="highlight">心</span><span class="highlight">形</span><span class="">岛</span></div>
              </div>
            </div>
            <div class="sug-item">

              <div>
                <div><span class="highlight">心</span><span class="highlight">形</span><span class="">脸</span><span class="">发</span><span class="">型</span><span class="">男</span></div>
              </div>
            </div>
            <div class="sug-item">

              <div>
                <div><span class="highlight">心</span><span class="highlight">形</span><span class="">树</span></div>
              </div>
            </div>
            <div class="sug-item">

              <div>
                <div><span class="highlight">心</span><span class="highlight">形</span><span class="">水</span><span class="">果</span><span class="">礼</span><span class="">盒</span></div>
              </div>
            </div>
            <div class="sug-item">

              <div>
                <div><span class="highlight">心</span><span class="highlight">形</span><span class="">慕</span><span class="">斯</span></div>
              </div>
            </div>
          </div>
        </div>
        <div>

        </div>
      </div>
    </div>
    <!-- 搜索区域2 -->
    <!-- <div class="sug-container-wrapper sug-pad">
      <div class="sug-container">
        <div class="history">
          <div class="header"><span> 历史记录 </span>
            <div class="icon-group">
              <div class="icon-box"><svg class="reds-icon icon"
                     width="16"
                     height="16">
                  <use xlink:href="#delete"></use>
                </svg>

              </div>

            </div>
          </div>
          <div class="history-list">
            <div class="history-item">兴东一房一厅</div>
            <div class="history-item">西丽租房子一室一厅</div>
            <div class="history-item">桌搭 mac</div>
            <div class="history-item">桌搭</div>
            <div class="history-item">京东京造K3</div>
            <div class="history-item">耳机</div>
            <div class="history-item">小红书网页版</div>

          </div>
        </div>
        <div class="sug-box">
          <div class="header">猜你想搜</div>
          <div class="sug-wrapper">
            <div class="sug-item query-trending">
              <div class="sug-text">macbookair劝退</div>

            </div>
            <div class="sug-item query-trending">
              <div class="sug-text">结肠炎能自愈吗</div>
            </div>
            <div class="sug-item query-trending">
              <div class="sug-text">兴东一房一厅</div>
            </div>
            <div class="sug-item query-trending">
              <div class="sug-text">实胖虚胖
              </div>
            </div>
            <div class="sug-item query-trending">
              <div class="sug-text">西丽租房子两房
              </div>
            </div>
            <div class="sug-item query-trending">
              <div class="sug-text">macbookpro必买配件
              </div>
            </div>
            <div class="sug-item query-trending">
              <div class="sug-text">咖啡控必入咖啡杯
              </div>
            </div>
            <div class="sug-item query-trending">
              <div class="sug-text">颞下颌关节弹响
              </div>
            </div>
          </div>
        </div>
        <div>
        </div>
      </div>
    </div> -->
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

// 清空搜索框
function clearSearch () {
  searchValue.value = ""
}

// 搜索关键词
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
		right: 10px;
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
	.sug_pad{
		// 弹出效果
		position: fixed;  // 使元素固定在视口中，不随页面滚动而移动。
		top: calc($layout-header-height - 24px);
		border-radius: $input-border-radius;
		background-color: $layout-body-background;
		// 显示在顶层
		z-index: 999;
		width: 100%;
		// 毛边效果
		box-shadow: 1px 2px 4px 1px rgba(0, 0, 0, 0.2);
		padding: 2px;
		.sug-container{
			.sug-box{
				.sug-wrapper{
					padding: 2px;
					.active{
							background-color: $active-color;
					}
					.sug-item{
						border-radius: $input-border-radius;
						padding: 8px 16px;
						cursor: pointer;
						&:hover{
							background-color: $active-color;
						}
						div {
							.highlight{
								color: $primary-color-red;
								font-weight: 500;
							}
							span{
								color: rgba(51, 51, 51, 0.6);
							}
						}
					}
				}
			}
		}
	}
}
</style>


