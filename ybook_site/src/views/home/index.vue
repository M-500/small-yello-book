<script setup lang="ts">
import tab_item from "@/components/TabItem/index.vue";
import note_item from "@/components/NoteItem/index.vue";
</script>

<template>
  <div class="container">
    <div class="channel_container">
      <tab_item name="美妆"
                :active="true" />
      <tab_item name="健身"
                :active="false" />
      <tab_item name="游泳"
                :active="false" />
      <tab_item name="育儿"
                :active="false" />
    </div>

    <div class="note_container">
      <note_item />
      <!-- <note_item style="width: 221.33333333333334px; padding-bottom: 16px; left: blur(42.5px); transform: translate(0px, 0px);" />
          <note_item />
          <note_item /> -->
    </div>

  </div>
</template>

<style scoped lang="scss">  
.container {
  .channel_container{
			overflow: hidden;
			display: flex;
			user-select: none;
			-webkit-user-select: none;
			align-items: center;
			font-size: 16px;
			color: var(--color-secondary-label);
			height: 40px;
			white-space: nowrap;
			margin-bottom: 16px;
		}
		.note_container{
			display: flex;
		}
}
</style>