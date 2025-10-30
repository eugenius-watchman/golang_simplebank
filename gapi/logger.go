package gapi

import (
	"context"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


func GrpcLogger(
	ctx context.Context, 
	req any, 
	info *grpc.UnaryServerInfo, 
	handler grpc.UnaryHandler,
) (resp any, err error) {
	startTime := time.Now()
	result, err := handler(ctx, req)
	duration := time.Since(startTime)

	statusCode := codes.Unknown
	if st, ok := status.FromError(err); ok {
		statusCode = st.Code()
	}

	logger := log.Info()
	if err != nil {
		logger = log.Error().Err(err)
	}

	logger.Str("protocol", "grpc").
		Str("method", info.FullMethod).
		Int("status_code", int(statusCode)).
		Str("status_text", statusCode.String()).
		Dur("duration", duration).
		Msg("received a gRPC request")

	return result, err
}

type ResponseRecorder struct {
	http.ResponseWriter
	StatusCode 	int
	Body		[]byte
}

func (rec *ResponseRecorder) WriteHeader(statusCode int) {
	rec.StatusCode = statusCode
	rec.ResponseWriter.WriteHeader(statusCode)
}

func (rec *ResponseRecorder) Write(body []byte) (int, error) {
	rec.Body = body
	return rec.ResponseWriter.Write(body)
}

// http logger middleware func
func HTTPLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		startTime := time.Now()
		rec := &ResponseRecorder{
			ResponseWriter: res,
			StatusCode: 	http.StatusOK,
		}

		handler.ServeHTTP(rec, req)
		duration := time.Since(startTime)

		logger := log.Info()
		if rec.StatusCode != http.StatusOK {
			logger = log.Error().Bytes("body", rec.Body)
		}

		logger.Str("protocol", "http").
			Str("method", req.Method).
			Str("path", req.RequestURI).
			Int("status_code", rec.StatusCode).
			Str("status_text", http.StatusText(rec.StatusCode)).
			Dur("duration", duration).
			Msg("received a HTTP request")
	})
}



// import (
// 	"context"
// 	"net/http"
// 	"time"

// 	"github.com/rs/zerolog/log"

// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/codes"
// 	"google.golang.org/grpc/status"
// )

// // GrpcLogger is a gRPC unary server interceptor that logs incoming requests
// // It captures method name, status code, duration, and any errors
// func GrpcLogger(
// 	ctx context.Context, 
// 	req any, 
// 	info *grpc.UnaryServerInfo, 
// 	handler grpc.UnaryHandler,
// ) (resp any, err error) {
// 	// Record start time to calculate request duration
// 	startTime := time.Now()
	
// 	// Call the actual gRPC handler
// 	result, err := handler(ctx, req)
	
// 	// Calculate how long the request took
// 	duration := time.Since(startTime)

// 	// Extract gRPC status code from error (default to Unknown if no error)
// 	statusCode := codes.Unknown
// 	if st, ok := status.FromError(err); ok {
// 		statusCode = st.Code()
// 	}

// 	// Choose appropriate log level based on whether there was an error
// 	logger := log.Info()
// 	if err != nil {
// 		logger = log.Error().Err(err) // Include error details if request failed
// 	}

// 	// Log structured information about the gRPC request
// 	logger.Str("protocol", "grpc").
// 		Str("method", info.FullMethod).        // Full gRPC method name (e.g., /package.Service/Method)
// 		Int("status_code", int(statusCode)).   // gRPC status code as integer
// 		Str("status_text", statusCode.String()). // gRPC status code as string (e.g., "OK", "NOT_FOUND")
// 		Dur("duration", duration).              // Request duration for performance monitoring
// 		Msg("received a gRPC request")         // Log message

// 	return result, err
// }

// // ResponseRecorder wraps http.ResponseWriter to capture response details
// // This allows us to log the status code and body of HTTP responses
// type ResponseRecorder struct {
// 	http.ResponseWriter
// 	StatusCode int     // Captured HTTP status code (e.g., 200, 404, 500)
// 	Body       []byte  // Captured response body (for error responses)
// }

// // WriteHeader overrides the default to capture the status code
// func (rec *ResponseRecorder) WriteHeader(statusCode int) {
// 	rec.StatusCode = statusCode
// 	rec.ResponseWriter.WriteHeader(statusCode)
// }

// // Write overrides the default to capture the response body
// func (rec *ResponseRecorder) Write(body []byte) (int, error) {
// 	rec.Body = body
// 	return rec.ResponseWriter.Write(body)
// }

// // HTTPLogger is a middleware that logs HTTP requests and responses
// // It wraps the HTTP handler to capture request details, status codes, and duration
// func HTTPLogger(handler http.Handler) http.Handler {
// 	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
// 		// Record start time for duration calculation
// 		startTime := time.Now()
		
// 		// Create a response recorder to capture status code and body
// 		rec := &ResponseRecorder{
// 			ResponseWriter: res,
// 			StatusCode:     http.StatusOK, // Default to 200 OK
// 		}

// 		// Pass the request through the handler chain
// 		handler.ServeHTTP(rec, req)
		
// 		// Calculate request duration
// 		duration := time.Since(startTime)

// 		// Choose appropriate log level based on HTTP status code
// 		logger := log.Info()
// 		if rec.StatusCode != http.StatusOK {
// 			// For non-200 responses, log as error and include response body
// 			logger = log.Error().Bytes("body", rec.Body)
// 		}

// 		// Log structured information about the HTTP request
// 		logger.Str("protocol", "http").
// 			Str("method", req.Method).                    // HTTP method (GET, POST, etc.)
// 			Str("path", req.RequestURI).                  // Request path/URL
// 			Int("status_code", rec.StatusCode).           // HTTP status code (200, 404, etc.)
// 			Str("status_text", http.StatusText(rec.StatusCode)). // Status text ("OK", "Not Found")
// 			Dur("duration", duration).                    // Request duration
// 			Msg("received a HTTP request")                // Log message
// 	})
// }