<template>
	<view class="content">
		<view class="input-group">
			<view class="input-row border">
				<m-input type="text" focus clearable v-model="mobile" placeholder="请输入手机号码"></m-input>
			</view>
			<view class="input-row">
				<m-input type="password" displayable v-model="pwd" placeholder="请输入密码"></m-input>
			</view>
		</view>
		<view class="btn-row">
			<button type="primary" class="primary" @tap="register">注册</button>
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
			}
		},
		methods: {
			register() {
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
				helper.ajax.post('/v1/vis/register', data, res => {
					if (res.status === 200) {
						helper.setSession(res.data,ret=>{
							uni.showToast({
								title: '注册成功'
							});
							setTimeout(() => {
								uni.switchTab({
									url: '../tabbar/tabbar-1/tabbar-1'
								})
							}, 1500)
						})
						return
					}
					console.log(res.message)
				})
			}
		}
	}
</script>

<style lang="scss">
	.input-group{
		margin: 0 20upx;
		m-input{
			min-height: 84upx;
		}
	}
</style>
