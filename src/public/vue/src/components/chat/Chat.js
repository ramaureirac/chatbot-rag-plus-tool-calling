import { onMounted } from 'vue'

export default {
  name: 'Chat',
  data() {
    return {
      xAnonID: '',
      user: 'user',
      lock: false,
      newMsg: '',
      messages: [],
    }
  },
  methods: {
    async authenticate() {
      try {
        const response = await fetch('http://localhost:8080/login', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
        })
        const data = await response.json()
        if (data.message) {
          this.xAnonID = data.message
          console.log('ok: successfully connected to backend api')
        } else {
          console.error('err: unable to retrive auth token')
        }
      } catch (err) {
        console.error('err: unable to communicate with backend api')
      }
    },

    async postMessage(text) {
      return fetch('http://localhost:8080/ask', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'X-Anon-ID': this.xAnonID,
        },
        body: JSON.stringify({
          message: text,
        }),
      })
    },

    async sendMessage() {
      if (this.lock) returnn
      const text = this.newMsg.trim()
      if (!text) return
      this.lock = true
      this.messages.push({ user: this.user, text: text })
      this.newMsg = ''
      this.scrollToEnd()
      try {
        let response = await this.postMessage(text)
        if (response.status == 403) {
          await this.authenticate()
          response = await this.postMessage(text)
        }
        const body = await response.json()
        this.messages.push({ user: 'assistant', text: body.message })
        this.scrollToEnd()
      } catch(e) {
        console.error('err: unable to communicate with backend api' + e)
      }
      this.lock = false
    },

    scrollToEnd() {
      this.$nextTick(() => {
        const el = this.$refs.messagesEnd
        el.scrollTop = el.scrollHeight
      })
    },
  },
}
