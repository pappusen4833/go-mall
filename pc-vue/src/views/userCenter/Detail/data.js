export default {
    buttonlist: [
        // 已完成条件： refundStatus === 2 || 已收货
        // 取消订单状态： paid===0 (未支付) && payTime > createTime + 15分钟
        // 未支付：10 拼团：11 拼团失败：12 已完成: 4
        // -1 : 申请退款
        // -2 : 退货成功
        // 0：待发货
        // 1：待收货
        // 2：已收货
        // 3：待评价
        // -1：已退款
        {
            'des': '取消订单',
            'status': [0, 10],
            'type': 'normal',
            'interface': 'cancelOrder'
        },
        {
            'des': '立即付款',
            'status': [10],
            'type': 'special',
            'interface': 'payNow'
        },
        // {
        //     'des': '申请退款',
        //     'status': [-1, 0, 11],
        //     'type': 'normal',
        //     'interface': 'getPayBack'
        // },
        {
            'des': '邀请好友拼团',
            'status': [11],
            'type': 'special',
            'interface': 'inviCombination'
        },
        // {
        //     'des': '取消退款',
        //     'status': [-1],
        //     'type': 'normal',
        //     'interface': 'cancelPayBack'
        // },
        {
            'des': '删除订单',
            'status': [-2, 2, 4, 12],
            'type': 'normal',
            'interface': 'deleteOrder'
        },
        {
            'des': '再次开团',
            'status': [12],
            'type': 'special',
            'interface': 'againCombination'
        },
        // {
        //     'des': '提醒发货',
        //     'status': [0],
        //     'type': 'special',
        //     'interface': 'remindShip'
        // },
   
        {
            'des': '确认收货',
            'status': [1],
            'type': 'special',
            'interface': 'sureReceive'
        },
        {
            'des': '评价',
            'status': [2, 3],
            'type': 'normal',
            'interface': 'subAssess'
        },
        {
            'des': '再次购买',
            'status': [3, 4],
            'type': 'special',
            'interface': 'buyAgain'
        }
    ]
}
