<template>
  <div ref="chartRef" style="height: 500px;"></div>
</template>

<script setup lang="ts">
import * as echarts from 'echarts';
import imag1 from '@/assets/a.png';
import {onMounted, ref} from "vue";

const chartRef = ref()


const createChart = () => {
  const chart = echarts.init(chartRef.value);

  const option: any = {
    title: {
      text: "",
    },
    tooltip: {},
    animationDurationUpdate: 1500,
    animationEasingUpdate: "quinticInOut",
    label: {
      normal: {
        show: true,
        textStyle: {
          fontSize: 12,
        },
      },
    },
    legend: {
      x: "center",
      show: false,
      data: ["夫妻", "战友", "亲戚"],
    },
    series: [
      {
        type: "graph",
        layout: "force",
        symbolSize: 45,
        focusNodeAdjacency: true,
        roam: true,
        categories: [
          {
            name: "夫妻",
            itemStyle: {
              normal: {
                color: "#009800",
              },
            },
          },
          {
            name: "战友",
            itemStyle: {
              normal: {
                color: "#4592FF",
              },
            },
          },
          {
            name: "亲戚",
            itemStyle: {
              normal: {
                color: "#3592F",
              },
            },
          },
        ],
        label: {
          show: true,
          position: 'bottom',
          distance: 5,
          color: '#333',
          fontFamily: 'sans-serif',
          fontSize: 14,
          fontWeight: 'bold',
          formatter: '{b}',
        },
        force: {
          repulsion: 1000,
        },
        edgeSymbolSize: [4, 50],
        edgeLabel: {
          normal: {
            show: true,
            textStyle: {
              fontSize: 10,
            },
            formatter: "{c}",
          },
        },
        data: [
          {
            id: "111",
            name: "徐贱云",
            draggable: true,
            symbol: `image://${imag1}`,
            symbolSize: [40, 40],
            itemStyle: {
              normal: {
                borderColor: '#fff',
                borderWidth: 2,
                borderType: 'solid',
                shadowColor: '#ccc',
                shadowBlur: 5,
                shadowOffsetX: 0,
                shadowOffsetY: 5,
                color: '#fff',
              },
            }
          },
          {
            name: "冯可梁",
            category: 1,
            draggable: true,
          },
          {
            name: "邓志荣",
            category: 1,
            draggable: true,
          },
          {
            name: "李荣庆",
            category: 1,
            draggable: true,
          },
          {
            name: "郑志勇",
            category: 1,
            draggable: true,
          },
          {
            name: "赵英杰",
            category: 1,
            draggable: true,
          },
          {
            name: "王承军",
            category: 1,
            draggable: true,
          },
          {
            name: "陈卫东",
            category: 1,
            draggable: true,
          },
          {
            name: "邹劲松",
            category: 1,
            draggable: true,
          },
          {
            name: "赵成",
            category: 1,
            draggable: true,
          },
          {
            name: "陈现忠",
            category: 1,
            draggable: true,
          },
          {
            name: "陶泳",
            category: 1,
            draggable: true,
          },
          {
            name: "王德福",
            category: 1,
            draggable: true,
          },
        ],
        links: [
          {
            source: 0,
            target: 1,
            category: 0,
            value: "夫妻",
          },
          {
            source: 0,
            target: 2,
            value: "子女",
          },
          {
            source: 0,
            target: 3,
            value: "夫妻",
          },
          {
            source: 0,
            target: 4,
            value: "父母",
          },
          {
            source: 0,
            target: 5,
            value: "朋友",
          },
          {
            source: 1,
            target: 2,
            value: "表亲",
          },
          {
            source: 4,
            target: 5,
            value: "朋友",
          },
          {
            source: 2,
            target: 8,
            value: "叔叔",
          },
          {
            source: 0,
            target: 12,
            value: "朋友",
          },
          {
            source: 6,
            target: 11,
            value: "爱人",
          },
          {
            source: 6,
            target: 3,
            value: "朋友",
          },
          {
            source: 7,
            target: 5,
            value: "朋友",
          },
          {
            source: 9,
            target: 10,
            value: "朋友",
          },
          {
            source: 3,
            target: 10,
            value: "朋友",
          },
          {
            source: 2,
            target: 11,
            value: "同学",
          },
        ],
        lineStyle: {
          normal: {
            opacity: 0.9,
            width: 1,
            curveness: 0,
          },
        },
      },
    ],
  };


  chart.setOption(option);
  chart.on('click', function (params) {
    console.log(params);
  });
};


onMounted(() => {
  createChart();
});


</script>

<style scoped lang="scss">
.echarts-sym-image > img {
  width: 100%;
  height: 100%;
  border-radius: 50%;
}
</style>