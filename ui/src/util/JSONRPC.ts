interface Request {
  resolve(info: any): void
  reject(err: any): void
}

export default class JSONRPC {
  private i: number
  private requests: (Request | null)[]
  public ws: WebSocket

  constructor(addr: string) {
    this.i = 0
    this.requests = []
    this.ws = new WebSocket(addr);
    this.ws.onmessage = (e: MessageEvent) => {
      this.onClientMessage(e)
    }
  }

  onClientMessage(event: MessageEvent) {
    let ret = JSON.parse(event.data)
    if (this.requests[ret.id] == null) {
      return 
    }
    const req = this.requests[ret.id] as Request
    if (ret.error == null) {
      req.resolve(ret.result);
    } else {
      req.reject(ret.error);
    };
    this.requests[ret.id] = null;
  }

  call<P, R>(method: string, param: P) : Promise<R> {
    return new Promise<R>((resolve, reject) => {
      while( this.requests[this.i] != null ) {
        this.i++
      }
      const id = this.i
      const params = [param]
      const request = {method, params, id}
      this.requests[id] = {
        resolve: (info: any) => resolve(info as R),
        reject
      }
      this.ws.send(JSON.stringify(request))
    })
  }
}

