<template>
	<view class="content">
		全部作业中的项目（以及按车牌号查找历史作业）
		<view id="btn-order" @tap="goToPage('../../order/order')">
			<text class="iconfont icon-baoyang"></text>
		</view>
	</view>
</template>

<script>
	import helper from '@/common/js/helper.js'
	export default {
		data() {
			return {
				title: 'Hello'
			}
		},
		onShow() {
			helper.initSession()
		},
		methods: {
			goToPage(url) {
				if (!url) return;
				uni.navigateTo({
					url
				});
			}
		}
	}
</script>

<style lang="scss">
#btn-order{
	position: fixed;
	bottom: 40rpx;
	right: 20rpx;
	width: 100rpx;
	height: 100rpx;
	line-height: 100rpx;
	text-align: center;
	font-size: 50rpx;
	background: #007AFF;
	color: #FFFFFF;
	border-radius: 50%;
}
</style>
