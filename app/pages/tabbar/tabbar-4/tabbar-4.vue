<template>
	<view class="content">
		消息 - 目前不知道有什么用
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
