<template>
	<view class="content">
		全部作业中的项目（以及按车牌号查找历史作业）
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

		}
	}
</script>

<style>

</style>
