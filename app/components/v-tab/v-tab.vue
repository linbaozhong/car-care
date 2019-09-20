<template>
	<view class="v-tab">
		<view :class="index===selectIndex?'action':''" v-for="(item,index) in items" :key="index" @click="onSelect(index)">
			{{item.name}}
		</view>
	</view>
</template>

<script>
	export default {
		name: "vTab",
		data() {
			return {
				selectIndex: 0,
			};
		},
		props: ['items'],
		methods: {
			onSelect(idx) {
				this.selectIndex = idx
				this.$emit('select', this.items[idx])
			}
		}
	}
</script>

<style lang="scss">
	.v-tab {
		height: 80upx;
		line-height: 80upx;
		color: $uni-text-color;
		background: white;
		box-shadow: 0 2upx 5upx $uni-bg-color-grey;
		display: flex;
		justify-content: space-around;
		align-items: stretch;
		font-size: $uni-font-size-base;

		.action {
			font-weight: bold;
			border-bottom: 4upx solid $uni-color-primary;
		}
	}
</style>
