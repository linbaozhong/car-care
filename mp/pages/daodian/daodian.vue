<template>
	<view id="daodian">
		<view class="records" v-for="record in records" :key="record.id">
			<view class="header">
				<view class="logo">
					<image :src="record.car.brand.logo_url" mode="aspectFit"></image>
					<text>{{record.car.brand.name}}</text>
				</view>
				<text class="gray">{{dateFormat(record.into_time)}} 到店</text>
			</view>
			<view class="item" @click="onViewReport(1,record.preview.id)">
				<view class="left">
					<text class="iconfont icon-yujian"></text>
					<view class="">
						<view>
							预检单
						</view>
						<view class="gray">
							{{dateFormat(record.preview.qc_time)}}
						</view>
					</view>
				</view>
				<view class="right">
					<view class="iconfont icon-you"></view>
					<uni-badge :text="'异常('+record.preview.abnormal+')'" type="warning" />
				</view>
			</view>
			<view class="item" @click="onViewReport(2,record.online_view.id)">
				<view class="left">
					<text class="iconfont icon-shangxianjiancha"></text>
					<view class="">
						<view>
							上线检查单
						</view>
						<view class="gray">
							{{dateFormat(record.online_view.qc_time)}}
						</view>
					</view>
				</view>
				<view class="right">
					<view class="iconfont icon-you"></view>
					<uni-badge :text="'异常('+record.online_view.abnormal+')'" type="warning" />
				</view>
			</view>
			<view class="item" @click="onViewReport(3,record.service_report.id)">
				<view class="left">
					<text class="iconfont icon-shigongdan"></text>
					<view class="">
						<view>
							施工报告
						</view>
						<view class="gray">
							{{dateFormat(record.service_report.rpt_time)}}
						</view>
					</view>
				</view>
				<view class="right">
					<view class="iconfont icon-you"></view>
				</view>
			</view>
			<view class="item" @click="onViewReport(4,record.qc_report.id)">
				<view class="left">
					<text class="iconfont icon-zhijiandan"></text>
					<view class="">
						<view>
							质检报告
						</view>
						<view class="gray">
							{{dateFormat(record.qc_report.rpt_time)}}
						</view>
					</view>
				</view>
				<view class="right">
					<view class="iconfont icon-you"></view>
				</view>
			</view>
		</view>
	</view>
</template>

<script>
	import uniBadge from '@/components/uni-badge/uni-badge.vue'
	import {
		dateUtils
	} from '../../common/js/util.js'
	export default {
		components: {
			uniBadge
		},
		data() {
			return {
				records: [{
					id: 1,
					car: {
						id: 1,
						plate: '京N9CD10',
						brand: {
							name: '昊锐',
							logo_url: 'https://car3.autoimg.cn/cardfs/series/g29/M0B/AE/D4/100x100_f40_autohomecar__wKgHG1s9t3CAQfVMAABHgh0VMdw516.png',
							maker: '上海大众',
						},
						model: '手自一体 优雅版',
						years: '2012',
						odometer: 5600,
					},
					odometer: 5600,
					into_time: 1568780293449,
					preview: {
						id: 1,
						qc_time: 1568780293449,
						abnormal: 2,
					},
					online_view: {
						id: 1,
						qc_time: 1568780293449,
						abnormal: 1,
					},
					service_report: {
						id: 1,
						rpt_time: 1568780293449,
					},
					qc_report: {
						id: 1,
						rpt_time: 1568780293449,
					},
				}, {
					id: 1,
					car: {
						id: 1,
						plate: '京N9CD10',
						brand: {
							name: '昊锐',
							logo_url: 'https://car3.autoimg.cn/cardfs/series/g29/M0B/AE/D4/100x100_f40_autohomecar__wKgHG1s9t3CAQfVMAABHgh0VMdw516.png',
							maker: '上海大众',
						},
						model: '手自一体 优雅版',
						years: '2012',
					odometer: 5600,
					},
					odometer: 5600,
					into_time: 1568780293449,
					preview: {
						id: 1,
						qc_time: 1568780293449,
						abnormal: 2,
					},
					online_view: {
						id: 1,
						qc_time: 1568780293449,
						abnormal: 1,
					},
					service_report: {
						id: 1,
						rpt_time: 1568780293449,
					},
					qc_report: {
						id: 1,
						rpt_time: 1568780293449,
					},
				}]
			}
		},
		computed: {
			dateFormat() {
				return ms => {
					return dateUtils.format(ms)
				}
			}
		},
		methods: {
			onViewReport(type, id) {
				uni.navigateTo({
					url: 'report/report?type=' + type + '&id=' + id
				})
			}
		}
	}
</script>

<style lang="scss">
	#daodian {
		.records {
			background: #FFFFFF;
			margin-bottom: 20upx;

			.gray {
				color: #CCCCCC;
			}

			.header {
				padding: 20upx;
				display: flex;
				justify-content: space-between;
				align-items: center;
				border-bottom: 1px solid $uni-border-color;

				.logo {
					display: flex;
					align-items: center;
					font-size: 32upx;

					image {
						width: 80upx;
						height: 80upx;
					}
				}

				text {
					margin-left: 10upx;
				}
			}

			.item {
				padding: 10upx 30upx;
				display: flex;
				justify-content: space-between;
				align-items: center;

				.left,
				.right {
					display: flex;
					align-items: center;

					view {
						margin: 0 10upx;
					}
				}

				.left .iconfont {
					width: 55upx;
					height: 55upx;
					font-size: 32upx;
					border: 1px solid #CCCCCC;
					border-radius: 50%;
					text-align: center;
				}

				.right {
					flex-direction: row-reverse;
				}

				.gray {
					font-size: 24upx;
				}
			}
		}
	}
</style>
