<template>
  <div class="container">
    <div class="img">
      <div class="media">
        <img src="@/assets/imgs/curry.jpeg"
             alt="">
        <div class="mask"></div>
        <div class="play-permission">
          <svg-icon class="per-icon"
                    color="#fff"
                    name="lock" />
          <span class="per-title">仅自己可见</span>
        </div>
      </div>
    </div>
    <div class="info">
      <div class="box">
        <div class="title">辞职啦</div>
        <div class="icon-list">
          <div class="icon">
            <svg-icon name="eye"
                      class="svg-ic"
                      color="#999"
                      hight="16px"
                      width="16px" />
            <span>0</span>
          </div>
          <div class="icon">
            <svg-icon name="comment"
                      class="svg-ic"
                      hight="16px"
                      color="#999"
                      width="16px" />
            <span>0</span>
          </div>
          <div class="icon">
            <svg-icon name="love"
                      class="svg-ic"
                      hight="16px"
                      color="#999"
                      width="16px" />
            <span>0</span>
          </div>
          <div class="icon">
            <svg-icon name="collect"
                      class="svg-ic"
                      hight="16px"
                      color="#999"
                      width="16px" />
            <span>0</span>
          </div>
          <div class="icon">
            <svg-icon name="share"
                      class="svg-ic"
                      hight="16px"
                      color="#999"
                      width="16px" />
            <span>0</span>
          </div>
        </div>
        <div class="time-status">
          <div class="status">已发布</div>
          <div class="time">2022年09月17日 22:09</div>
        </div>
      </div>
      <div class="control-list">
        <div class="control">
          <svg-icon name="permission"
                    class="svg-ic1"
                    hight="16px"
                    width="16px" />
          <span class="title">权限设置</span>
        </div>
        <div class="control">
          <svg-icon name="cancel-top"
                    class="svg-ic1"
                    hight="16px"
                    width="16px" />
          <span class="title">取消置顶</span>
        </div>
        <div class="control">
          <svg-icon name="edit"
                    class="svg-ic1"
                    hight="16px"
                    width="16px" />
          <span class="title">编辑</span>
        </div>
        <div class="control">
          <svg-icon name="delete"
                    class="svg-ic1"
                    hight="16px"
                    width="16px" />
          <span class="title">删除</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { } from "vue"
</script>

<style lang="scss" scoped>
.container{
  height: 192px;
  padding: 16px;
  border-radius: 4px;
  text-decoration: none;
  display: flex;
  -webkit-user-select: none;
  user-select: none;
  .img{
    height: 160px;
    .media{
      width: 120px;
      height: 160px;
      position: relative;
      img{
        width: 100%;
        height: 160px;
        object-fit: cover;
        border-radius: 4px;
      }
      .mask[data-v-5223e286]:before {
        content: "";
        position: absolute;
        left: 0;
        right: 0;
        bottom: 0;
        top: 0;
        background-color: #000;
        opacity: .3;
        z-index: 2;
        border-radius: 4px;
      }
      .play-permission{
        font-size: 11px;
        border-radius: 12px;
        left: 4px;
        position: absolute;
        padding: 2px 4px;
        color: #fff;
        // background: rgba(0, 0, 0, .2);
        bottom: 8px;
        z-index: 5;
        display: flex;
        align-content: center;
        .per-icon{
          width: 16px;
          height: 16px;
          margin-right: 2px;
        }
        .per-title{
          font-size: 12px;
          color: #fff;
          align-items: center;
        }
      }
    }
  }
  .info{
    margin-left: 16px;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    margin-top: 1px;
    .box{
      .title{
        color: rgba(0, 0, 0, .82);
        font-weight: 500;
        line-height: 20px;
        font-family: PingFang SC;
      }
      .icon-list{
        margin-top: 8px;
        display: flex;
        align-items: center;
        align-content: center;
        .icon{
          color: #999;
          display: flex;
          align-items: center;
          .svg-ic{
            width: 16px;
            height: 16px;
            vertical-align: text-bottom;
            color: #999;
            // margin-right: 8px;
          }
          span{
            margin-left: 5px;
            font-size: 12px;
            color: #999;
            margin-right: 15px;
            align-items: center;
          }

        }
      }
      .time-status{
        margin-top: 8px;
        display: flex;  
        .status{
          color: #42c9a0;
          font-size: 12px;
          line-height: 20px;
        }
        .time{
          color: #999;
          margin-left: 12px;
          font-size: 12px;
          line-height: 20px;
        }
      }
    }
    .control-list{
      display: flex;
      align-items: center;
      .control{
        // visibility: hidden;
        display: flex;
        align-items: center;
        margin-right: 5px;
        cursor: pointer;
        .svg-ic1{

        }
        .title{
          font-size: 14px;
          line-height: 16px;
          color: #666;
          margin-left: 4px;
        }
      }
      .control:hover{
        .svg-ic1{
          color: $primary-buttom-active-colder;
        }
        .title{
          color: $primary-buttom-active-colder;
          font-weight: 600;
        }
      }
    }
  }
}
</style>
