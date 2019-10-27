<template>
	<view class="content">
		<view class="">
			我的个人信息
		</view>
		<view class="">
			我的作业统计（本月，上月），收入统计
		</view>
		<view class="">
			店长独享：1、门店信息维护；2、店员信息维护
		</view>
		<view class="list">
			<view class="item">
				<view class="">
					<text class="iconfont icon-guanyu"></text>
					<text class="">关于黄白格</text>
				</view>
				<view class="iconfont icon-you">

				</view>
			</view>
			<view class="item">
				<view class="">
					<text class="iconfont icon-kefu"></text>
					<text class="">客服帮助</text>
				</view>
				<view class="iconfont icon-you">

				</view>
			</view>
		</view>
		<view class="list">
				<view class="item" @click="onLogout">
					<view class="">
						<text class="iconfont icon-guanyu"></text>
						<text class="">退出</text>
					</view>
					<view class="iconfont icon-you"></view>
				</view>
			</view>
		
	</view>
</template>

<script>
	import helper from '@/common/js/helper.js'

	export default {
		data() {
			return {
			}
		},
		onShow() {
helper.initSession()
		},
		methods: {
			onLogout(){
				helper.ajax.post('/v1/clerk/logout',{},res=>{
					if(res.status===200){
						helper.delSession()
						uni.reLaunch({
							url:'../../login/login'
						})
						return
					}
					console.log(res.message);
				})
			}
		}
	}
</script>

<style lang="scss">
	.list {
		background: #FFFFFF;
		margin: 20rpx 0;
		padding-left: 30rpx;

		.item {
			padding: 20rpx 20rpx 20rpx 10rpx;
			display: flex;
			justify-content: space-between;
			align-items: center;
			border-bottom: 1px solid $uni-border-color;

			text {
				margin-right: 20rpx;
			}
		}
	}
</style>
