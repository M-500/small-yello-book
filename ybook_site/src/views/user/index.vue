<template>
  <div class="container">
    <div class="user_info">
      <div class="user">
        <div class="avatar">
          <div class="avatar_wrap">
            <img :src="userDetail.avatar"
                 alt="">
          </div>
        </div>
        <div class="info">
          <div class="info_part">
            <div class="info">
              <div class="basic_info">
                <div class="nickname">
                  <div class="user_name">{{ userDetail.nickName }}</div>
                </div>
                <div class="user_content">
                  <span class="yellow_id">小黄书号：{{ userDetail.globalNumber }}</span>
                  <span class="user_ip">IP归属地：上海</span>
                </div>

              </div>
              <div class="user_desc">{{ userDetail.signature }}</div>
              <div class="user_tag">
                <div class="tag_item">
                  <SvgIcon icon="icon-tag"
                           class="icon"
                           name="man"
                           width="12"
                           high="12"
                           color="#333" />
                  <span class="gener_text">27岁</span>
                </div>
              </div>
              <div class="data_info">
                <div class="box">
                  <span class="count">14</span>
                  <span class="show">关注</span>
                </div>
                <div class="box">
                  <span class="count">6</span>
                  <span class="show">粉丝</span>
                </div>
                <div class="box">
                  <span class="count">14</span>
                  <span class="show">获赞与收藏</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="tab_list">
      <tab_item :name="tab.label"
                v-for="(tab,index) in tabs"
                :key="index"
                :class="{active: selectTab === tab.name}"
                @click="selectTab = tab.name" />
    </div>
    <div class="feeds-tab-container">
      <note-item v-for="(item,index) in userNoteList"
                 :key="index"
                 :item="item" />
    </div>
  </div>
</template>

<script  setup>
import NoteItem from "@/components/note-item/index.vue";
import tab_item from "@/components/TabItem/index.vue";
import { useRoute } from 'vue-router';
import { getUserInfoByUUID } from '@/api/user';
import { ref } from "vue";
import { getUserNoteListRequest } from '@/api/note/index'
import { onMounted } from 'vue'
// import { queryNoteListForm } from '@/api/note/types'

const userNoteList = ref([])

// 获取用户对应的笔记列表
const getNoteListData = async () => {
  const params = {
    state: -1,
    page: 1,
    size: 10
  }
  getUserNoteListRequest(params).then((res) => {
    console.log("叼你个嗨", res.data.list)
    userNoteList.value = res.data.list
  });
}

const selectTab = ref('note');
const tabs = ref([
  { name: 'note', label: '笔记' },
  { name: 'collect', label: '收藏' },
  { name: 'like', label: '点赞' },
]);
const route = useRoute();
const userId = route.params.uuid; // 获取路由参数上的uuid
const userDetail = ref({});

// 获取用户信息
const getUserInfo = async () => {
  const res = await getUserInfoByUUID(userId);
  userDetail.value = res.data;
};

onMounted(() => {
  getUserInfo();
  getNoteListData();
});

</script>

<style lang="scss" scoped>
.container{
  .user_info{
    // padding-top: 72px;
    display: flex;
    align-items: center;
    justify-content: center;
  }
  .user{
    display: flex;
    width: 100%;
    padding: 24px 0;
    justify-content: center;
    margin: 0 120px;
    .avatar{
      display: flex;
      align-content: center;  // 居中对齐
      .avatar_wrap{
        width: calc(25vw - 30px); // 25vw代表屏幕宽度的25%
        height: calc(17.5vw - 21px); //
        display: flex;
        justify-content: center;
        align-items: center;
        img{
          border-radius: 50%;
          margin: 0 auto;
          width: 70%;
          height: 100%;
          object-fit: cover;
        }
      }
    }
    .info{
      position: relative;
      width: 100%;
      .info_part{
      margin-left: 24px;
      .basic_info{
        display: flex;
        flex-direction: column;
        width: 100%;
        .nickname{
          font-weight: 600;
          font-size: 24px;
          line-height: 120%;
          color: #333;
          word-wrap: break-word;
          width: 100%;
        }
        .user_content{
          width: 100%;
          font-size: 12px;
          line-height: 120%;
          color: var(--color-tertiary-label);
          display: flex;
          margin-top: 8px;
        }
      }
      .user_desc{
        width: 100%;
        font-size: 14px;
        line-height: 140%;
        color: #333;
        margin-top: 16px;
        white-space: pre-line;
      }
      .user_tag{
        margin-top: 16px;
        display: flex;
        align-items: center;
        font-size: 12px;
        color: #333;
        text-align: center;
        font-weight: 400;
        line-height: 120%;
        flex-wrap: wrap;
        grid-row-gap: var(--verticalGapPx);
        row-gap: 12px;
        .tag_item{
          margin-right: 0;
          display: flex;
          justify-content: flex-start;
          align-content: center;
          span{
            margin-left: 4px;
            color: rgba(51,51,51,0.6);
          }
        }
      }
      .data_info{
        display: flex;
        align-items: center;
        justify-content: flex-start;
        margin-top: 20px;
        .box{
          eight: 100%;
          display: flex;
          align-items: center;
          justify-content: center;
          text-align: center;
          margin-right: 16px;
          .count{
            font-weight: 500;
            font-size: 14px;
            margin-right: 4px;
          }
          .show{
            color: rgba(51, 51, 51, 0.6);
            font-size: 14px;
            line-height: 120%;
          }
        }
      }
    }
  }
    
  }
  .tab_list{
    display: flex;
    justify-content: center;
  }
  .feeds-tab-container{
    display: flex; // flex布局
    flex-wrap: wrap; // 自动换行
    justify-content: flex-start; // 两端对齐
    gap: 20px;  // 间距
  }
}
</style>
