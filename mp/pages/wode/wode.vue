<template>
	<view id="wode">
		<view class="avatar">
			<image :src="userInfo.avatarUrl" mode="aspectFit"></image>
			<view class="nickName">
				<text v-if="userInfo.nickName" class="nickName">{{userInfo.nickName}}</text>
				<button v-else type="primary" plain="true" open-type="getUserInfo" size="mini" @getuserinfo="getUserInfo">获取昵称</button>
				<!-- <text class="iconfont icon-you"></text> -->
			</view>
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
	</view>
</template>

<script>
	import helper from '@/common/js/helper.js'
	export default {
		data() {
			return {
				userInfo: {
					avatarUrl: '../../static/logo.png',
					nickName: '',
				} //getApp().globalData.session,
			};
		},
		onLoad() {
			uni.checkSession({
				fail() {
					helper.login()
				}
			})
		},
		methods: {
			getPhoneNumber(res) {
				if (res.detail.errMsg === 'getPhoneNumber:ok') {
					helper.ajax.post('/v1/openuser/bindmobilefromweixin', {
						data: res.detail.encryptedData,
						iv: res.detail.iv,
					}, e => {
						if (e.success) {
							uni.showToast({
								title: e.data
							})
							helper.setSession({
								mobile: e.data
							}, res => {
								this.userInfo = helper.getSession()
							})
						} else {
							uni.showToast({
								title: e.message,
								icon: 'none'
							})
						}
					})
				} else {
					uni.navigateTo({
						url: 'bindMobile'
					})
				}
			},
			getUserInfo(res) {
				if (res.detail.errMsg === 'getUserInfo:ok') {
					helper.setSession(res.detail.userInfo, res => {
						this.userInfo = helper.getSession()
					})
				} else {
					uni.showToast({
						title: res.detail.errMsg
					})
				}
			},
		}
	}
</script>

<style lang="scss">
	#wode {
		.avatar {
			width: 750upx;
			height: 300upx;
			display: flex;
			flex-direction: column;
			justify-content: center;
			align-items: center;
			color: #FFFFFF;
			background:#DD4A68;
			
			image{
				width: 100upx;
				height: 100upx;
				border-radius: 50%;
			}
			.nickName {
				margin-top: 20upx;
				display:flex;
				justify-content: center;
				align-items: center;
				
				button {
					padding: 0;
					border: none;
					line-height: 1;
					color: inherit;
				}

				.iconfont {
					margin-left: 10upx;
				}
			}
		}
	
		.list{
			background: #FFFFFF;
			margin: 20upx 0;
			padding-left:30upx;
			
			.item{
				padding: 20upx 20upx 20upx 10upx;
				display:flex;
				justify-content: space-between;
				align-items: center;
				border-bottom:1px solid $uni-border-color;
				
				text{
					margin-right: 20upx;
				}
			}
		}
	}
</style>
