<template>
  <div ref="chartRef" style="height: 500px;"></div>
</template>

<script>
import * as echarts from 'echarts';
import img1 from '@/assets/a.png';
import img2 from '@/assets/b.png';
import { ref, onMounted } from 'vue';

export default {
  setup() {
    const chartRef = ref(null);

    const nodes = [
      { id: 1, name: 'Tom', img: img1 },
      { id: 2, name: 'Jerry', img: img2 },
      { id: 3, name: 'Mike', img: img1 },
      { id: 4, name: 'Bob', img: img2 },
    ];

    const links = [
      { source: 1, target: 2 },
      { source: 1, target: 3 },
      { source: 2, target: 4 },
      { source: 3, target: 4 },
    ];

    const createChart = () => {
      const chart = echarts.init(chartRef.value);

      const option = {
        tooltip: {
          show: true,
        },
        series: [
          {
            type: 'graph',
            layout: 'force',
            roam: true,
            symbolSize: [80, 80],
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
            edgeSymbol: ['none', 'arrow'],
            edgeSymbolSize: 10,
            data: nodes.map(node => {
              return {
                id: node.id,
                name: node.name,
                symbol: `image://${node.img}`,
                symbolSize: [60, 60], // 将高宽设为相同的值
                textStyle:{
                  backgroundColor:{image:img1},
                  borderRadius: 100,
                },
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
                },
              };
            }),
            links: links,
          },
        ],
      };

      chart.setOption(option);
    };

    onMounted(() => {
      createChart();
    });

    return {
      chartRef,
    };
  },
};
</script>

<style scoped>
.echarts-sym-image > img {
  width: 60px;
  height: 60px;
  border-radius: 50%;
}
</style>
