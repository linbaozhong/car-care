<template>
	<view>
		<view class="uni-margin-wrap">
			<swiper class="swiper" :indicator-dots="indicatorDots" :autoplay="autoplay" :interval="interval" :duration="duration">
				<swiper-item>
					<view class="swiper-item">750px * 350px</view>
				</swiper-item>
				<swiper-item>
					<view class="swiper-item">750px * 350px</view>
				</swiper-item>
				<swiper-item>
					<view class="swiper-item">750px * 350px</view>
				</swiper-item>
			</swiper>
		</view>
		<view class="content">
			<view class="header">{{store.name}}</view>
			<view class="">
				<uni-badge :text="store.tag" type="primary" />
				<text class="tag">总评分</text><text class="tag count">{{store.rating}}</text>
				<text class="tag">总订单</text><text class="tag count">{{store.orders}}</text>
			</view>
			<view class="address" @click="onGotoMap">
				<view class="">{{store.address}}</view>
				<view class="">12.22km <text class="iconfont icon-daohang"></text></view>
			</view>
		</view>
		<view class="services">
			<view class="header">
				服务项目
			</view>
			<view class="items">
				<view class="item">
					<view class="iconfont icon-wifi">
					</view>
					<view>连接WiFi</view>
				</view>
				<view class="item">
					<view class="iconfont icon-youhuiquan">
					</view>
					<view>领优惠券</view>
				</view>
				<view class="item">
					<view class="iconfont icon-xiche">
					</view>
					<view>洗车</view>
				</view>
				<view class="item">
					<view class="iconfont icon-baoyang">
					</view>
					<view>保养</view>
				</view>
			</view>
		</view>
	</view>
</template>

<script>
	import uniBadge from '@/components/uni-badge/uni-badge.vue'
	export default {
		components: {
			uniBadge
		},
		data() {
			return {
				store: {
					name: '黄白格汽车美容(亦庄店)',
					tag: '旗舰店',
					rating: 5,
					orders: 2816,
					address: '北京亦庄',
					longitude: 116.534686,
					latitude: 39.811717,
				},
				items: [

				],
				indicatorDots: false,
				autoplay: true,
				interval: 2000,
				duration: 500
			}
		},
		onLoad() {},
		methods: {
			onGotoMap() {
				uni.openLocation({
					name: this.store.name,
					address: this.store.address,
					latitude: this.store.latitude,
					longitude: this.store.longitude,
					success: function() {
						console.log('success');
					}
				})
			},
		}
	}
</script>

<style lang="scss">
	.uni-margin-wrap {
		width: 750upx;
		margin: 0;
		$swiper-height: 350upx;

		.swiper {
			height: $swiper-height;
		}

		.swiper-item {
			display: block;
			height: $swiper-height;
			line-height: $swiper-height;
			text-align: center;
			background: #007AFF;
		}
	}

	.content {
		padding: 20upx;
		background: #FFFFFF;
		font-size: 24upx;
		line-height: 60upx;

		.header {
			font-size: 36upx;
		}

		.tag {
			margin-left: 10upx;
		}

		.count {
			color: $uni-color-error;
		}

		.address {
			display: flex;
			direction: column;
			justify-content: space-between;
			align-items: center;

			text {
				margin-left: 10upx;
			}
		}
	}

	.services {
		margin: 20upx 0;
		padding: 20upx;
		background: #FFFFFF;
		.header{
			font-size: 32upx;
		}
		.items{
			margin: 20upx 0;
			display: flex;
			justify-content: space-around;
			align-items: center;
			
			.item{
				display: flex;
				flex-direction: column;
				align-items: center;
				justify-content:center;
				width: 20%;
				
				.iconfont{
					width: 80upx;
					height: 80upx;
					line-height: 80upx;
					border-radius: 50%;
					font-size: 56upx;
					color: #FFFFFF;
					background: #FFC000;
					text-align: center;
				}
			}
		}
	}
</style>
