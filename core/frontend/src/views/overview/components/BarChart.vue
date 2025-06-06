<template>
	<bt-charts :options="chartOptions" />
</template>

<script setup lang="ts">
import { useThemeVars } from 'naive-ui'
import { formatTime } from '@/utils'

const theme = useThemeVars()

const { dateType, chartData } = defineProps({
	dateType: {
		type: String,
		required: true,
	},
	chartData: {
		type: Object as PropType<{
			fail: {
				label: string
				value: [string, number, string][]
			}
			success: {
				label: string
				value: [string, number, string][]
			}
		}>,
		required: true,
	},
})

const chartOptions = computed(() => {
	const { success, fail } = chartData
	return {
		grid: {
			top: '16%',
			left: '2%',
			right: '2%',
			bottom: '4%',
			containLabel: true,
		},
		legend: {
			top: 0,
			itemGap: 16,
			icon: 'circle',
			itemWidth: 10,
			itemHeight: 10,
			textStyle: {
				color: theme.value.textColor2,
			},
			data: [success.label, fail.label],
		},
		tooltip: {
			trigger: 'axis',
			order: 'seriesDesc',
			axisPointer: {
				type: 'shadow',
			},
		},
		xAxis: {
			type: 'category',
			axisLabel: {
				formatter: (val: string) => {
					if (dateType === 'hourly') {
						return formatTime(val, 'HH:mm')
					}
					return formatTime(val, 'yyyy-MM-dd')
				},
			},
		},
		yAxis: {
			type: 'value',
			boundaryGap: [0, '6%'],
			splitLine: {
				show: true,
				lineStyle: {
					type: 'dashed',
					width: 1,
					color: theme.value.borderColor,
				},
			},
		},
		series: [
			{
				type: 'bar',
				name: fail.label,
				data: fail.value,
				stack: 'total',
				itemStyle: { color: '#1A519B' },
			},
			{
				type: 'bar',
				name: success.label,
				data: success.value,
				stack: 'total',
				itemStyle: { color: '#91CC75' },
			},
		],
	}
})
</script>
