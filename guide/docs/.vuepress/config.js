module.exports = {
  base: '/tank-websocket-go-server/',
  title: 'Tank WebSocket',
  description: 'A lightweight, feature-rich WebSocket server implementation in Go',
  themeConfig: {
    nav: [
      { text: 'Home', link: '/' },
      { text: 'Guide', link: '/guide/' },
      { text: 'GitHub', link: 'https://github.com/fanqie/tank-websocket-go-server' }
    ],
    sidebar: {
      '/guide/': [
        {
          title: 'Guide',
          collapsable: false,
          children: [
            '',
            'installation',
            'quick-start',
            'client-connection',
            'heartbeat',
            'topic-subscription',
            'authentication',
            'debug-logging'
          ]
        }
      ]
    }
  },
  locales: {
    '/': {
      lang: 'en-US',
      title: 'Tank WebSocket',
      description: 'A lightweight, feature-rich WebSocket server implementation in Go',
      selectText: 'Languages',
      label: 'English',
      ariaLabel: 'Languages',
      editLinkText: 'Edit this page on GitHub',
      lastUpdated: 'Last Updated',
      nav: [
        { text: 'Home', link: '/' },
        { text: 'Guide', link: '/guide/' },
        { text: 'GitHub', link: 'https://github.com/fanqie/tank-websocket-go-server' }
      ]
    },
    '/zh/': {
      lang: 'zh-CN',
      title: 'Tank WebSocket',
      description: '一个用 Go 语言实现的轻量级、功能丰富的 WebSocket 服务器',
      selectText: '选择语言',
      label: '简体中文',
      ariaLabel: '语言',
      editLinkText: '在 GitHub 上编辑此页',
      lastUpdated: '最后更新时间',
      nav: [
        { text: '首页', link: '/zh/' },
        { text: '指南', link: '/zh/guide/' },
        { text: 'GitHub', link: 'https://github.com/fanqie/tank-websocket-go-server' }
      ],
      sidebar: {
        '/zh/guide/': [
          {
            title: '指南',
            collapsable: false,
            children: [
              '',
              'installation',
              'quick-start',
              'client-connection',
              'heartbeat',
              'topic-subscription',
              'authentication',
              'debug-logging'
            ]
          }
        ]
      }
    }
  },
  plugins: [
    '@vuepress/back-to-top',
    '@vuepress/medium-zoom'
  ]
} 