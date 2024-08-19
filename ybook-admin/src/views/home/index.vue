<template>
  <div class="container">

    <el-card>
      <div class="user-info">
        <div class="avatar-box">
          <img src="../../assets/imgs/curry.jpeg"
               alt="">
        </div>
        <div class="info-box">
          <h4>巧舌如簧的哑巴</h4>
          <p class="detail">
            <span class="tips"><label for="">0</label>关注</span>
            <span class="tips"><label for="">1</label>粉丝</span>
            <span class="tips"><label for="">2</label>获赞与收藏</span>
            <span class="tips line"></span>
            <span class="tips">小红书号： wulinlin</span>
          </p>
          <p class="note">还没有简介</p>
        </div>
      </div>

    </el-card>

    <el-card class="data-container">
      <div class="data-view">
        <yb-title title="数据总览"></yb-title>
        <div class="block-line">
          <div class="block-container">
            <span class="des">新增粉丝</span>
            <span class="con">0</span>
          </div>
          <div class="block-container">
            <span class="des">主页访客</span>
            <span class="con">0</span>
          </div>
          <div class="block-container">
            <span class="des">观看数</span>
            <span class="con">0</span>
          </div>
          <div class="block-container">
            <span class="des">互动数</span>
            <span class="con">0</span>
          </div>
        </div>
        <yb-title title="最新发布笔记数据"></yb-title>
      </div>
    </el-card>
  </div>
</template>

<script setup>
import YbTitle from '@/components/YbTitle/index.vue'
</script>

<style lang="scss" scoped>
.container{
	height: 100%;
	.user-info{
		padding: 24px;
		display: flex;
		justify-content: flex-start;
		.avatar-box{
			width: 80px;
    	height: 80px;
			;
			img{
				height: 100%;
				vertical-align: middle;
				display: block;
				object-fit: cover;
				overflow-clip-margin: content-box;
				overflow: clip;
				border-radius: 100%;
				border: 1px solid rgb(232, 232, 232);
				line-height: 32px;
			}
		}
		.info-box{
			width: calc(100% - 80px);
			height: 80px;
			padding: 0px 24px;
			display: flex;
			flex-direction: column;
			justify-content: space-between;
			.detail{
				height: 20px;
				line-height: 20px;
				color: rgb(136, 136, 136);
				font-size: 14px;
				.tips{
					height: 20px;
					line-height: 20px;
					color: rgb(136, 136, 136);
					font-size: 14px;
					margin-right: 32px;
					label{
						margin-right: 8px;
						color: rgb(45, 45, 45);
						font-weight: 500;
					}
				}
				.line{
					margin-top: 2px;
					border-right: 1px solid rgb(232, 232, 232);
				}
			}
			.note{
				color: rgb(136, 136, 136);
			}
		}
	}
	.data-container{
		margin-top: 20px;
		height: calc(100% - 150px);
	}
	.data-view{
		padding: 0 24px;
		display: flex;
		flex-direction: column;
		.title-container{
			padding: 24px 0;
			.title-box{
				display: flex;
				align-items: center;
				float: left;
				.icon{
					display: inline-block;
					width: 5px;
					height: 18px;
					margin-right: 8px;
					border-radius: 16px;
					background: rgb(58, 100, 255);
				}
				.title-name{
					font-size: 20px;
					font-weight: 500;
					color: rgb(45, 45, 45);
					padding-right: 9px;
				}
			}
		}
		.block-line{
			display: flex;
    	justify-content: left;
			box-sizing: border-box;
			.block-container{
				width: 195px;
				height: 78px;
				padding: 16px;
				margin-right: 16px;
				border-radius: 4px;
				background: rgb(251, 251, 251);
				font-family: Avenir;
				.des{
					display: block;
					color: rgb(136, 136, 136);
					font-size: 12px;
					line-height: 12px;
				}
				.con{
					display: block;
					font-size: 22px;
					line-height: 22px;
					font-weight: 800;
					margin-top: 12px;
					color: rgb(45, 45, 45);
				}
			}
		}
	}
}

::v-deep(.el-card){
	border-radius: 8px;
}
::v-deep(.el-card__body){
	padding: 0px;
}
</style>
