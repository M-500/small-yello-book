<template>
    <div class="card_container">
        <div class="cover">
            <img src="https://th.bing.com/th/id/R.c8dc540e2aca707e4b6ce7dfc683df89?rik=6oYnoUTZcBLzOQ&riu=http%3a%2f%2fseopic.699pic.com%2fphoto%2f50053%2f4738.jpg_wh1200.jpg&ehk=ySgvy2yVdn%2bgvOxx24%2bX65oDzRtujqpdpNjrBPr23OM%3d&risl=&pid=ImgRaw&r=0" alt="">
        </div>
        <div class="bottom">
            <div class="title">一定要收藏</div>
            <div class="info">
                <div class="left">
                    <img src="https://th.bing.com/th/id/R.c8dc540e2aca707e4b6ce7dfc683df89?rik=6oYnoUTZcBLzOQ&riu=http%3a%2f%2fseopic.699pic.com%2fphoto%2f50053%2f4738.jpg_wh1200.jpg&ehk=ySgvy2yVdn%2bgvOxx24%2bX65oDzRtujqpdpNjrBPr23OM%3d&risl=&pid=ImgRaw&r=0" alt="">
                    <span>小明</span>
                </div>
                <div class="right">
                    <img src="../assets/icon/like.svg" alt="">
                    <span>1千+</span>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">

</script>

<style scoped>
.card_container{
    display: flex;
    flex-direction: column; /* 设置为上下布局 */
    justify-content: flex-start; /* 从上到下排列 */
    align-items: center;
    user-select: none;
    -webkit-user-select: none;
    width: 250px;
}

.cover{
    width: 100%;
    height: 300px;
    overflow: hidden;
    border-radius: 8px;
    margin-bottom: 16px;
}
.bottom{
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    width: 100%;
}
.info{
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    margin-top: 3px;
}

.left{
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;
}
.left img{
    width: 24px;
    height: 24px;
    border-radius: 50%;
}
.left span{
    margin-left: 4px;
    font-size: 13px;
}
.right{
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;
}
.right span{
    margin-left: 4px;
    font-size: 13px;
}
.right img{
    width: 16px;
}
</style>