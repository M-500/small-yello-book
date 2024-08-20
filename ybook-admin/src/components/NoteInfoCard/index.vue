<template>
  <div class="container">
    <div class="info-box">
      <div class="cover-page">

      </div>
      <div class="info-text">
        <span class="title">发财啦！</span>
        <span class="publish-time">发布于2022-01-01 19:21</span>
      </div>
    </div>
    <div class="info-list">
      <ul class="data-list">
        <li class="active">
          <span class="circle"></span>
          <label for="">观看量</label>
          <b class="align-text">379</b>
        </li>
        <li class="">
          <span class="circle"></span>
          <label for="">点赞量</label>
          <b class="align-text">379</b>
        </li>
        <li class="">
          <span class="circle"></span>
          <label for="">收藏数</label>
          <b class="align-text">379</b>
        </li>
      </ul>
      <ul class="data-list">
        <li class="">
          <span class="circle"></span>
          <label for="">评论数</label>
          <b class="align-text">379</b>
        </li>
        <li class="">
          <span class="circle"></span>
          <label for="">分享数</label>
          <b class="align-text">379</b>
        </li>
      </ul>
    </div>
  </div>
</template>

<script setup>
import { } from "vue"
</script>

<style lang="scss" scoped>
.container{
	height: 373px;
	padding: 20px;
	margin-right: 17px;
	margin-bottom: 20px;
	background: rgb(255, 255, 255);
	box-shadow: rgba(204, 204, 204, 0.2) 0px 4px 30px;
	border-radius: 8px;
	.info-box{
		width: 100%;
    height: 100px;
    border-bottom: 0.5px solid rgb(232, 232, 232);
    font-size: 12px;
    display: flex;
    justify-content: left;
		.cover-page{
			position: relative;
			width: 84px;
			height: 84px;
			text-align: center;
			border-radius: 4px;
			cursor: pointer;
			background: url('@/assets/imgs/curry.jpeg') 50% center / cover no-repeat scroll content-box content-box grey;
		}
		.info-text{
			overflow-wrap: break-word;
			max-width: 236px;
			padding-left: 12px;
			.title{
				font-size: 14px;
				line-height: 22px;
				max-height: 44px;
				font-weight: 500;
				display: -webkit-box;
				-webkit-box-orient: vertical;
				-webkit-line-clamp: 2;
				overflow: hidden;
				text-overflow: ellipsis;
			}
			.publish-time{
				display: block;
				line-height: 12px;
				margin-top: 12px;
				color: rgb(204, 204, 204);
			}
		}
	}
	.info-list{
		width: 430px;
    height: 137px;
    // margin-top: 16px;
    display: flex;
    justify-content: left;
		ul {
				display: block;
				list-style-type: disc;
				margin-block-start: 1em;
				margin-block-end: 1em;
				margin-inline-start: 0px;
				margin-inline-end: 0px;
				padding-inline-start: 40px;
				unicode-bidi: isolate;
		}
		.data-list{
			margin-left: 20px;
			display: inline-block;
			width: 148px;
			list-style: none;
			padding: 0px;
			margin: 0px;
			li{
				list-style: none;
				position: relative;
				line-height: 12px;
				white-space: nowrap;
				overflow: hidden;
				text-overflow: ellipsis;
				margin: 16px 0px 0px;
				// display: list-item;
				// text-align: -webkit-match-parent;
				// unicode-bidi: isolate;
			}
			.active{
				color: rgb(58, 100, 255);
				.circle{
					background-color: rgb(58, 100, 255);
				}
				label{
					color: rgb(58, 100, 255);
				}
				b{
					font-weight: 600;
				}
			}
			label{
				font-size: 12px;
				color: rgb(136, 136, 136);
				margin-left: 8px;
				margin-right: 16px;
				cursor: pointer;
			}
			b{
				margin-left: 20px;
				font-size: 12px;
				line-height: 12px;
				font-family: Avenir;
				white-space: nowrap;
				overflow: hidden;
				text-overflow: ellipsis;
				cursor: pointer;
			}
			.circle{
				display: block;
				width: 4px;
				height: 4px;
				border-radius: 4px;
				// background-color: rgb(58, 100, 255);
				position: absolute;
				top: 6px;
			}
		}
	}
}
</style>
