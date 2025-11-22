# Error handling

Go uses explicit error handling with error return values. Functions return error as last value. Create errors with ``errors.New()`` or ``fmt.Errorf()``. No exceptions, errors are values to be handled explicitly.

The error type is an interface type. An error variable represents any value that can describe itself as a string. This Go's error handling phylosophy provides consistent error representation across all Go code.

The ``errors.New()`` is the simplest way to create error values by taking a string message and returning an error implementing the error interface. Useful for simple, static error messages.

The ``fmt.Errorf()`` creates formatted error messages using printf-style verbs.

## Wrapping/Unwrapping Errors

Create error chains preserving original errors while adding context using ``fmt.Errorf()`` with ``%w`` verb. Use ``errors.Unwrap()``, ``errors.Is()``, and ``errors.As()`` to work with wrapped errors. Enables rich error context for easier debugging.

## panic and recover

``panic()`` stops execution and unwinds stack, ``recover()`` catches panics in deferred functions. Use sparingly for unrecoverable errors. While Go emphasizes explicit error, panic/recover serve as safety net for exceptional situations