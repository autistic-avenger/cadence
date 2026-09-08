import { NextURL } from 'next/dist/server/web/next-url'
import { cookies } from 'next/headers'
import { NextResponse } from 'next/server'
import type { NextRequest } from 'next/server'
 
export async function proxy(request: NextRequest) {
    const cookie = await cookies()
    const url =new URL(request.url)
    const jwt = cookie.get("token")
    if (jwt?.value == null && url.pathname != "/login"){
        return NextResponse.redirect(new URL('/login',request.url))
    }
    if (jwt?.value != null && url.pathname == "/login"){
        return NextResponse.redirect(new URL('/playground',request.url))
    }
    return 
}
 
export const config = {
  matcher: ['/playground','/login'],
}