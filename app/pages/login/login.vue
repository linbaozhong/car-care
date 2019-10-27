<template>
	<view class="content">
		<view class="input-group">
			<view class="input-row border">
				<m-input class="m-input" type="text" clearable focus v-model="mobile" placeholder="请输入手机号码"></m-input>
			</view>
			<view class="input-row">
				<m-input type="password" displayable v-model="pwd" placeholder="请输入密码"></m-input>
			</view>
		</view>
		<view class="btn-row">
			<button type="primary" class="primary" @tap="bindLogin">登录</button>
		</view>
		<view class="action-row">
			<navigator url="../register/register">注册账号</navigator>
			<text>|</text>
			<navigator url="../pwd/pwd">忘记密码</navigator>
		</view>
	</view>
</template>

<script>
	import helper from '@/common/js/helper.js';
	import util from '@/common/js/util.js';
	import mInput from '@/components/m-input.vue';
	export default {
		components: {
			mInput
		},
		data() {
			return {
				mobile: '',
				pwd: '',
			};
		},
		onLoad() {
			helper.initSession()
			if (!helper.APP_DATA.session) {
				return
			}
			helper.APP_DATA.session = helper.getSession()
			if (helper.APP_DATA.session) {
				uni.reLaunch({
					url: '../tabbar/tabbar-1/tabbar-1'
				})
				return
			}
		},
		methods: {
			bindLogin() {
				if (!util.isMobile(this.mobile)) {
					uni.showToast({
						icon: 'none',
						title: '手机号码为 11 个数字'
					});
					return;
				}
				if (this.pwd.length < 6 || this.pwd.length > 12) {
					uni.showToast({
						icon: 'none',
						title: '密码为 6~12 个字符'
					});
					return;
				}

				const data = {
					mobile: this.mobile,
					pwd: this.pwd,
				}
				helper.ajax.post('/v1/vis/login', data, res => {
					if (res.status === 200) {
						helper.setSession(res.data,ret=>{
							uni.showToast({
								title: '登录成功'
							});
							setTimeout(() => {
								uni.reLaunch({
									url: '../tabbar/tabbar-1/tabbar-1'
								})
							}, 1500)
						})
						return
					}
					uni.showToast({
						title:res.message,
						icon:'none',
					})
				})
			}
		}
	}
</script>

<style lang="scss">
	.input-group{
		margin: 0 20rpx;
		m-input{
			min-height: 84rpx;
		}
	}
	.action-row {
		display: flex;
		flex-direction: row;
		justify-content: center;
	}

	.action-row navigator {
		color: #007aff;
		padding: 0 20rpx;
	}

</style>
