# 'context' package

Carries deadlines, cancellation signals, and request-scoped values across API boundaries. Essential for robust concurrent applications, especially web services. Enables canelling long-running operations, setting timeouts, passing request data.

## Deadlines & Cancellations

Context package mechanisms for controlling operation lifetime and propagating cancellation signals. Supports deadlines or timeouts. Funcions should check 'ctx.Done()' and return early when cancelled. 

## Common uses

Common uses: 
- HTTP timeouts.
- Database deadlines.
- Goroutine canellation coordination.
- Request-scoped values.

Essential for web servers, microservices, circuit breakers and building responsive APIs that handle cancellation gracefully.