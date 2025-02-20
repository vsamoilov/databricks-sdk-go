// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// The Jobs API allows you to create, edit, and delete jobs.
package jobs

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/databricks/databricks-sdk-go/useragent"
)

type JobsInterface interface {
	// WithImpl could be used to override low-level API implementations for unit
	// testing purposes with [github.com/golang/mock] or other mocking frameworks.
	// Deprecated: use MockJobsInterface instead.
	WithImpl(impl JobsService) JobsInterface

	// Impl returns low-level Jobs API implementation
	// Deprecated: use MockJobsInterface instead.
	Impl() JobsService

	// WaitGetRunJobTerminatedOrSkipped repeatedly calls [JobsAPI.GetRun] and waits to reach TERMINATED or SKIPPED state
	WaitGetRunJobTerminatedOrSkipped(ctx context.Context, runId int64,
		timeout time.Duration, callback func(*Run)) (*Run, error)

	// Cancel all runs of a job.
	//
	// Cancels all active runs of a job. The runs are canceled asynchronously, so it
	// doesn't prevent new runs from being started.
	CancelAllRuns(ctx context.Context, request CancelAllRuns) error

	// Cancel a run.
	//
	// Cancels a job run or a task run. The run is canceled asynchronously, so it
	// may still be running when this request completes.
	CancelRun(ctx context.Context, cancelRun CancelRun) (*WaitGetRunJobTerminatedOrSkipped[struct{}], error)

	// Calls [JobsAPIInterface.CancelRun] and waits to reach TERMINATED or SKIPPED state
	//
	// You can override the default timeout of 20 minutes by calling adding
	// retries.Timeout[Run](60*time.Minute) functional option.
	//
	// Deprecated: use [JobsAPIInterface.CancelRun].Get() or [JobsAPIInterface.WaitGetRunJobTerminatedOrSkipped]
	CancelRunAndWait(ctx context.Context, cancelRun CancelRun, options ...retries.Option[Run]) (*Run, error)

	// Cancel a run.
	//
	// Cancels a job run or a task run. The run is canceled asynchronously, so it
	// may still be running when this request completes.
	CancelRunByRunId(ctx context.Context, runId int64) error

	CancelRunByRunIdAndWait(ctx context.Context, runId int64, options ...retries.Option[Run]) (*Run, error)

	// Create a new job.
	//
	// Create a new job.
	Create(ctx context.Context, request CreateJob) (*CreateResponse, error)

	// Delete a job.
	//
	// Deletes a job.
	Delete(ctx context.Context, request DeleteJob) error

	// Delete a job.
	//
	// Deletes a job.
	DeleteByJobId(ctx context.Context, jobId int64) error

	// Delete a job run.
	//
	// Deletes a non-active run. Returns an error if the run is active.
	DeleteRun(ctx context.Context, request DeleteRun) error

	// Delete a job run.
	//
	// Deletes a non-active run. Returns an error if the run is active.
	DeleteRunByRunId(ctx context.Context, runId int64) error

	// Export and retrieve a job run.
	//
	// Export and retrieve the job run task.
	ExportRun(ctx context.Context, request ExportRunRequest) (*ExportRunOutput, error)

	// Get a single job.
	//
	// Retrieves the details for a single job.
	Get(ctx context.Context, request GetJobRequest) (*Job, error)

	// Get a single job.
	//
	// Retrieves the details for a single job.
	GetByJobId(ctx context.Context, jobId int64) (*Job, error)

	// Get job permission levels.
	//
	// Gets the permission levels that a user can have on an object.
	GetPermissionLevels(ctx context.Context, request GetJobPermissionLevelsRequest) (*GetJobPermissionLevelsResponse, error)

	// Get job permission levels.
	//
	// Gets the permission levels that a user can have on an object.
	GetPermissionLevelsByJobId(ctx context.Context, jobId string) (*GetJobPermissionLevelsResponse, error)

	// Get job permissions.
	//
	// Gets the permissions of a job. Jobs can inherit permissions from their root
	// object.
	GetPermissions(ctx context.Context, request GetJobPermissionsRequest) (*JobPermissions, error)

	// Get job permissions.
	//
	// Gets the permissions of a job. Jobs can inherit permissions from their root
	// object.
	GetPermissionsByJobId(ctx context.Context, jobId string) (*JobPermissions, error)

	// Get a single job run.
	//
	// Retrieve the metadata of a run.
	GetRun(ctx context.Context, request GetRunRequest) (*Run, error)

	// Get the output for a single run.
	//
	// Retrieve the output and metadata of a single task run. When a notebook task
	// returns a value through the `dbutils.notebook.exit()` call, you can use this
	// endpoint to retrieve that value. Databricks restricts this API to returning
	// the first 5 MB of the output. To return a larger result, you can store job
	// results in a cloud storage service.
	//
	// This endpoint validates that the __run_id__ parameter is valid and returns an
	// HTTP status code 400 if the __run_id__ parameter is invalid. Runs are
	// automatically removed after 60 days. If you to want to reference them beyond
	// 60 days, you must save old run results before they expire.
	GetRunOutput(ctx context.Context, request GetRunOutputRequest) (*RunOutput, error)

	// Get the output for a single run.
	//
	// Retrieve the output and metadata of a single task run. When a notebook task
	// returns a value through the `dbutils.notebook.exit()` call, you can use this
	// endpoint to retrieve that value. Databricks restricts this API to returning
	// the first 5 MB of the output. To return a larger result, you can store job
	// results in a cloud storage service.
	//
	// This endpoint validates that the __run_id__ parameter is valid and returns an
	// HTTP status code 400 if the __run_id__ parameter is invalid. Runs are
	// automatically removed after 60 days. If you to want to reference them beyond
	// 60 days, you must save old run results before they expire.
	GetRunOutputByRunId(ctx context.Context, runId int64) (*RunOutput, error)

	// List jobs.
	//
	// Retrieves a list of jobs.
	//
	// This method is generated by Databricks SDK Code Generator.
	List(ctx context.Context, request ListJobsRequest) listing.Iterator[BaseJob]

	// List jobs.
	//
	// Retrieves a list of jobs.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListAll(ctx context.Context, request ListJobsRequest) ([]BaseJob, error)

	// BaseJobSettingsNameToJobIdMap calls [JobsAPI.ListAll] and creates a map of results with [BaseJob].Settings.Name as key and [BaseJob].JobId as value.
	//
	// Returns an error if there's more than one [BaseJob] with the same .Settings.Name.
	//
	// Note: All [BaseJob] instances are loaded into memory before creating a map.
	//
	// This method is generated by Databricks SDK Code Generator.
	BaseJobSettingsNameToJobIdMap(ctx context.Context, request ListJobsRequest) (map[string]int64, error)

	// GetBySettingsName calls [JobsAPI.BaseJobSettingsNameToJobIdMap] and returns a single [BaseJob].
	//
	// Returns an error if there's more than one [BaseJob] with the same .Settings.Name.
	//
	// Note: All [BaseJob] instances are loaded into memory before returning matching by name.
	//
	// This method is generated by Databricks SDK Code Generator.
	GetBySettingsName(ctx context.Context, name string) (*BaseJob, error)

	// List job runs.
	//
	// List runs in descending order by start time.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListRuns(ctx context.Context, request ListRunsRequest) listing.Iterator[BaseRun]

	// List job runs.
	//
	// List runs in descending order by start time.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListRunsAll(ctx context.Context, request ListRunsRequest) ([]BaseRun, error)

	// Repair a job run.
	//
	// Re-run one or more tasks. Tasks are re-run as part of the original job run.
	// They use the current job and task settings, and can be viewed in the history
	// for the original job run.
	RepairRun(ctx context.Context, repairRun RepairRun) (*WaitGetRunJobTerminatedOrSkipped[RepairRunResponse], error)

	// Calls [JobsAPIInterface.RepairRun] and waits to reach TERMINATED or SKIPPED state
	//
	// You can override the default timeout of 20 minutes by calling adding
	// retries.Timeout[Run](60*time.Minute) functional option.
	//
	// Deprecated: use [JobsAPIInterface.RepairRun].Get() or [JobsAPIInterface.WaitGetRunJobTerminatedOrSkipped]
	RepairRunAndWait(ctx context.Context, repairRun RepairRun, options ...retries.Option[Run]) (*Run, error)

	// Update all job settings (reset).
	//
	// Overwrite all settings for the given job. Use the [_Update_
	// endpoint](:method:jobs/update) to update job settings partially.
	Reset(ctx context.Context, request ResetJob) error

	// Trigger a new job run.
	//
	// Run a job and return the `run_id` of the triggered run.
	RunNow(ctx context.Context, runNow RunNow) (*WaitGetRunJobTerminatedOrSkipped[RunNowResponse], error)

	// Calls [JobsAPIInterface.RunNow] and waits to reach TERMINATED or SKIPPED state
	//
	// You can override the default timeout of 20 minutes by calling adding
	// retries.Timeout[Run](60*time.Minute) functional option.
	//
	// Deprecated: use [JobsAPIInterface.RunNow].Get() or [JobsAPIInterface.WaitGetRunJobTerminatedOrSkipped]
	RunNowAndWait(ctx context.Context, runNow RunNow, options ...retries.Option[Run]) (*Run, error)

	// Set job permissions.
	//
	// Sets permissions on a job. Jobs can inherit permissions from their root
	// object.
	SetPermissions(ctx context.Context, request JobPermissionsRequest) (*JobPermissions, error)

	// Create and trigger a one-time run.
	//
	// Submit a one-time run. This endpoint allows you to submit a workload directly
	// without creating a job. Runs submitted using this endpoint don’t display in
	// the UI. Use the `jobs/runs/get` API to check the run state after the job is
	// submitted.
	Submit(ctx context.Context, submitRun SubmitRun) (*WaitGetRunJobTerminatedOrSkipped[SubmitRunResponse], error)

	// Calls [JobsAPIInterface.Submit] and waits to reach TERMINATED or SKIPPED state
	//
	// You can override the default timeout of 20 minutes by calling adding
	// retries.Timeout[Run](60*time.Minute) functional option.
	//
	// Deprecated: use [JobsAPIInterface.Submit].Get() or [JobsAPIInterface.WaitGetRunJobTerminatedOrSkipped]
	SubmitAndWait(ctx context.Context, submitRun SubmitRun, options ...retries.Option[Run]) (*Run, error)

	// Update job settings partially.
	//
	// Add, update, or remove specific settings of an existing job. Use the [_Reset_
	// endpoint](:method:jobs/reset) to overwrite all job settings.
	Update(ctx context.Context, request UpdateJob) error

	// Update job permissions.
	//
	// Updates the permissions on a job. Jobs can inherit permissions from their
	// root object.
	UpdatePermissions(ctx context.Context, request JobPermissionsRequest) (*JobPermissions, error)
}

func NewJobs(client *client.DatabricksClient) *JobsAPI {
	return &JobsAPI{
		JobsService: &jobsImpl{
			client: client,
		},
	}
}

// The Jobs API allows you to create, edit, and delete jobs.
//
// You can use a Databricks job to run a data processing or data analysis task
// in a Databricks cluster with scalable resources. Your job can consist of a
// single task or can be a large, multi-task workflow with complex dependencies.
// Databricks manages the task orchestration, cluster management, monitoring,
// and error reporting for all of your jobs. You can run your jobs immediately
// or periodically through an easy-to-use scheduling system. You can implement
// job tasks using notebooks, JARS, Delta Live Tables pipelines, or Python,
// Scala, Spark submit, and Java applications.
//
// You should never hard code secrets or store them in plain text. Use the
// [Secrets CLI] to manage secrets in the [Databricks CLI]. Use the [Secrets
// utility] to reference secrets in notebooks and jobs.
//
// [Databricks CLI]: https://docs.databricks.com/dev-tools/cli/index.html
// [Secrets CLI]: https://docs.databricks.com/dev-tools/cli/secrets-cli.html
// [Secrets utility]: https://docs.databricks.com/dev-tools/databricks-utils.html#dbutils-secrets
type JobsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(JobsService)
	JobsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
// Deprecated: use MockJobsInterface instead.
func (a *JobsAPI) WithImpl(impl JobsService) JobsInterface {
	a.JobsService = impl
	return a
}

// Impl returns low-level Jobs API implementation
// Deprecated: use MockJobsInterface instead.
func (a *JobsAPI) Impl() JobsService {
	return a.JobsService
}

// WaitGetRunJobTerminatedOrSkipped repeatedly calls [JobsAPI.GetRun] and waits to reach TERMINATED or SKIPPED state
func (a *JobsAPI) WaitGetRunJobTerminatedOrSkipped(ctx context.Context, runId int64,
	timeout time.Duration, callback func(*Run)) (*Run, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	return retries.Poll[Run](ctx, timeout, func() (*Run, *retries.Err) {
		run, err := a.GetRun(ctx, GetRunRequest{
			RunId: runId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		if callback != nil {
			callback(run)
		}
		status := run.State.LifeCycleState
		statusMessage := fmt.Sprintf("current status: %s", status)
		if run.State != nil {
			statusMessage = run.State.StateMessage
		}
		switch status {
		case RunLifeCycleStateTerminated, RunLifeCycleStateSkipped: // target state
			return run, nil
		case RunLifeCycleStateInternalError:
			err := fmt.Errorf("failed to reach %s or %s, got %s: %s",
				RunLifeCycleStateTerminated, RunLifeCycleStateSkipped, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// WaitGetRunJobTerminatedOrSkipped is a wrapper that calls [JobsAPI.WaitGetRunJobTerminatedOrSkipped] and waits to reach TERMINATED or SKIPPED state.
type WaitGetRunJobTerminatedOrSkipped[R any] struct {
	Response *R
	RunId    int64 `json:"run_id"`
	Poll     func(time.Duration, func(*Run)) (*Run, error)
	callback func(*Run)
	timeout  time.Duration
}

// OnProgress invokes a callback every time it polls for the status update.
func (w *WaitGetRunJobTerminatedOrSkipped[R]) OnProgress(callback func(*Run)) *WaitGetRunJobTerminatedOrSkipped[R] {
	w.callback = callback
	return w
}

// Get the Run with the default timeout of 20 minutes.
func (w *WaitGetRunJobTerminatedOrSkipped[R]) Get() (*Run, error) {
	return w.Poll(w.timeout, w.callback)
}

// Get the Run with custom timeout.
func (w *WaitGetRunJobTerminatedOrSkipped[R]) GetWithTimeout(timeout time.Duration) (*Run, error) {
	return w.Poll(timeout, w.callback)
}

// Cancel a run.
//
// Cancels a job run or a task run. The run is canceled asynchronously, so it
// may still be running when this request completes.
func (a *JobsAPI) CancelRun(ctx context.Context, cancelRun CancelRun) (*WaitGetRunJobTerminatedOrSkipped[struct{}], error) {
	err := a.JobsService.CancelRun(ctx, cancelRun)
	if err != nil {
		return nil, err
	}
	return &WaitGetRunJobTerminatedOrSkipped[struct{}]{

		RunId: cancelRun.RunId,
		Poll: func(timeout time.Duration, callback func(*Run)) (*Run, error) {
			return a.WaitGetRunJobTerminatedOrSkipped(ctx, cancelRun.RunId, timeout, callback)
		},
		timeout:  20 * time.Minute,
		callback: nil,
	}, nil
}

// Calls [JobsAPI.CancelRun] and waits to reach TERMINATED or SKIPPED state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[Run](60*time.Minute) functional option.
//
// Deprecated: use [JobsAPI.CancelRun].Get() or [JobsAPI.WaitGetRunJobTerminatedOrSkipped]
func (a *JobsAPI) CancelRunAndWait(ctx context.Context, cancelRun CancelRun, options ...retries.Option[Run]) (*Run, error) {
	wait, err := a.CancelRun(ctx, cancelRun)
	if err != nil {
		return nil, err
	}
	tmp := &retries.Info[Run]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(tmp)
	}
	wait.timeout = tmp.Timeout
	wait.callback = func(info *Run) {
		for _, o := range options {
			o(&retries.Info[Run]{
				Info:    info,
				Timeout: wait.timeout,
			})
		}
	}
	return wait.Get()
}

// Cancel a run.
//
// Cancels a job run or a task run. The run is canceled asynchronously, so it
// may still be running when this request completes.
func (a *JobsAPI) CancelRunByRunId(ctx context.Context, runId int64) error {
	return a.JobsService.CancelRun(ctx, CancelRun{
		RunId: runId,
	})
}

func (a *JobsAPI) CancelRunByRunIdAndWait(ctx context.Context, runId int64, options ...retries.Option[Run]) (*Run, error) {
	return a.CancelRunAndWait(ctx, CancelRun{
		RunId: runId,
	}, options...)
}

// Delete a job.
//
// Deletes a job.
func (a *JobsAPI) DeleteByJobId(ctx context.Context, jobId int64) error {
	return a.JobsService.Delete(ctx, DeleteJob{
		JobId: jobId,
	})
}

// Delete a job run.
//
// Deletes a non-active run. Returns an error if the run is active.
func (a *JobsAPI) DeleteRunByRunId(ctx context.Context, runId int64) error {
	return a.JobsService.DeleteRun(ctx, DeleteRun{
		RunId: runId,
	})
}

// Get a single job.
//
// Retrieves the details for a single job.
func (a *JobsAPI) GetByJobId(ctx context.Context, jobId int64) (*Job, error) {
	return a.JobsService.Get(ctx, GetJobRequest{
		JobId: jobId,
	})
}

// Get job permission levels.
//
// Gets the permission levels that a user can have on an object.
func (a *JobsAPI) GetPermissionLevelsByJobId(ctx context.Context, jobId string) (*GetJobPermissionLevelsResponse, error) {
	return a.JobsService.GetPermissionLevels(ctx, GetJobPermissionLevelsRequest{
		JobId: jobId,
	})
}

// Get job permissions.
//
// Gets the permissions of a job. Jobs can inherit permissions from their root
// object.
func (a *JobsAPI) GetPermissionsByJobId(ctx context.Context, jobId string) (*JobPermissions, error) {
	return a.JobsService.GetPermissions(ctx, GetJobPermissionsRequest{
		JobId: jobId,
	})
}

// Get the output for a single run.
//
// Retrieve the output and metadata of a single task run. When a notebook task
// returns a value through the `dbutils.notebook.exit()` call, you can use this
// endpoint to retrieve that value. Databricks restricts this API to returning
// the first 5 MB of the output. To return a larger result, you can store job
// results in a cloud storage service.
//
// This endpoint validates that the __run_id__ parameter is valid and returns an
// HTTP status code 400 if the __run_id__ parameter is invalid. Runs are
// automatically removed after 60 days. If you to want to reference them beyond
// 60 days, you must save old run results before they expire.
func (a *JobsAPI) GetRunOutputByRunId(ctx context.Context, runId int64) (*RunOutput, error) {
	return a.JobsService.GetRunOutput(ctx, GetRunOutputRequest{
		RunId: runId,
	})
}

// List jobs.
//
// Retrieves a list of jobs.
//
// This method is generated by Databricks SDK Code Generator.
func (a *JobsAPI) List(ctx context.Context, request ListJobsRequest) listing.Iterator[BaseJob] {

	getNextPage := func(ctx context.Context, req ListJobsRequest) (*ListJobsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.JobsService.List(ctx, req)
	}
	getItems := func(resp *ListJobsResponse) []BaseJob {
		return resp.Jobs
	}
	getNextReq := func(resp *ListJobsResponse) *ListJobsRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// List jobs.
//
// Retrieves a list of jobs.
//
// This method is generated by Databricks SDK Code Generator.
func (a *JobsAPI) ListAll(ctx context.Context, request ListJobsRequest) ([]BaseJob, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[BaseJob](ctx, iterator)
}

// BaseJobSettingsNameToJobIdMap calls [JobsAPI.ListAll] and creates a map of results with [BaseJob].Settings.Name as key and [BaseJob].JobId as value.
//
// Returns an error if there's more than one [BaseJob] with the same .Settings.Name.
//
// Note: All [BaseJob] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *JobsAPI) BaseJobSettingsNameToJobIdMap(ctx context.Context, request ListJobsRequest) (map[string]int64, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "name-to-id")
	mapping := map[string]int64{}
	result, err := a.ListAll(ctx, request)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		key := v.Settings.Name
		_, duplicate := mapping[key]
		if duplicate {
			return nil, fmt.Errorf("duplicate .Settings.Name: %s", key)
		}
		mapping[key] = v.JobId
	}
	return mapping, nil
}

// GetBySettingsName calls [JobsAPI.BaseJobSettingsNameToJobIdMap] and returns a single [BaseJob].
//
// Returns an error if there's more than one [BaseJob] with the same .Settings.Name.
//
// Note: All [BaseJob] instances are loaded into memory before returning matching by name.
//
// This method is generated by Databricks SDK Code Generator.
func (a *JobsAPI) GetBySettingsName(ctx context.Context, name string) (*BaseJob, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "get-by-name")
	result, err := a.ListAll(ctx, ListJobsRequest{})
	if err != nil {
		return nil, err
	}
	tmp := map[string][]BaseJob{}
	for _, v := range result {
		key := v.Settings.Name
		tmp[key] = append(tmp[key], v)
	}
	alternatives, ok := tmp[name]
	if !ok || len(alternatives) == 0 {
		return nil, fmt.Errorf("BaseJob named '%s' does not exist", name)
	}
	if len(alternatives) > 1 {
		return nil, fmt.Errorf("there are %d instances of BaseJob named '%s'", len(alternatives), name)
	}
	return &alternatives[0], nil
}

// List job runs.
//
// List runs in descending order by start time.
//
// This method is generated by Databricks SDK Code Generator.
func (a *JobsAPI) ListRuns(ctx context.Context, request ListRunsRequest) listing.Iterator[BaseRun] {

	getNextPage := func(ctx context.Context, req ListRunsRequest) (*ListRunsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.JobsService.ListRuns(ctx, req)
	}
	getItems := func(resp *ListRunsResponse) []BaseRun {
		return resp.Runs
	}
	getNextReq := func(resp *ListRunsResponse) *ListRunsRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// List job runs.
//
// List runs in descending order by start time.
//
// This method is generated by Databricks SDK Code Generator.
func (a *JobsAPI) ListRunsAll(ctx context.Context, request ListRunsRequest) ([]BaseRun, error) {
	iterator := a.ListRuns(ctx, request)
	return listing.ToSlice[BaseRun](ctx, iterator)
}

// Repair a job run.
//
// Re-run one or more tasks. Tasks are re-run as part of the original job run.
// They use the current job and task settings, and can be viewed in the history
// for the original job run.
func (a *JobsAPI) RepairRun(ctx context.Context, repairRun RepairRun) (*WaitGetRunJobTerminatedOrSkipped[RepairRunResponse], error) {
	repairRunResponse, err := a.JobsService.RepairRun(ctx, repairRun)
	if err != nil {
		return nil, err
	}
	return &WaitGetRunJobTerminatedOrSkipped[RepairRunResponse]{
		Response: repairRunResponse,
		RunId:    repairRun.RunId,
		Poll: func(timeout time.Duration, callback func(*Run)) (*Run, error) {
			return a.WaitGetRunJobTerminatedOrSkipped(ctx, repairRun.RunId, timeout, callback)
		},
		timeout:  20 * time.Minute,
		callback: nil,
	}, nil
}

// Calls [JobsAPI.RepairRun] and waits to reach TERMINATED or SKIPPED state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[Run](60*time.Minute) functional option.
//
// Deprecated: use [JobsAPI.RepairRun].Get() or [JobsAPI.WaitGetRunJobTerminatedOrSkipped]
func (a *JobsAPI) RepairRunAndWait(ctx context.Context, repairRun RepairRun, options ...retries.Option[Run]) (*Run, error) {
	wait, err := a.RepairRun(ctx, repairRun)
	if err != nil {
		return nil, err
	}
	tmp := &retries.Info[Run]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(tmp)
	}
	wait.timeout = tmp.Timeout
	wait.callback = func(info *Run) {
		for _, o := range options {
			o(&retries.Info[Run]{
				Info:    info,
				Timeout: wait.timeout,
			})
		}
	}
	return wait.Get()
}

// Trigger a new job run.
//
// Run a job and return the `run_id` of the triggered run.
func (a *JobsAPI) RunNow(ctx context.Context, runNow RunNow) (*WaitGetRunJobTerminatedOrSkipped[RunNowResponse], error) {
	runNowResponse, err := a.JobsService.RunNow(ctx, runNow)
	if err != nil {
		return nil, err
	}
	return &WaitGetRunJobTerminatedOrSkipped[RunNowResponse]{
		Response: runNowResponse,
		RunId:    runNowResponse.RunId,
		Poll: func(timeout time.Duration, callback func(*Run)) (*Run, error) {
			return a.WaitGetRunJobTerminatedOrSkipped(ctx, runNowResponse.RunId, timeout, callback)
		},
		timeout:  20 * time.Minute,
		callback: nil,
	}, nil
}

// Calls [JobsAPI.RunNow] and waits to reach TERMINATED or SKIPPED state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[Run](60*time.Minute) functional option.
//
// Deprecated: use [JobsAPI.RunNow].Get() or [JobsAPI.WaitGetRunJobTerminatedOrSkipped]
func (a *JobsAPI) RunNowAndWait(ctx context.Context, runNow RunNow, options ...retries.Option[Run]) (*Run, error) {
	wait, err := a.RunNow(ctx, runNow)
	if err != nil {
		return nil, err
	}
	tmp := &retries.Info[Run]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(tmp)
	}
	wait.timeout = tmp.Timeout
	wait.callback = func(info *Run) {
		for _, o := range options {
			o(&retries.Info[Run]{
				Info:    info,
				Timeout: wait.timeout,
			})
		}
	}
	return wait.Get()
}

// Create and trigger a one-time run.
//
// Submit a one-time run. This endpoint allows you to submit a workload directly
// without creating a job. Runs submitted using this endpoint don’t display in
// the UI. Use the `jobs/runs/get` API to check the run state after the job is
// submitted.
func (a *JobsAPI) Submit(ctx context.Context, submitRun SubmitRun) (*WaitGetRunJobTerminatedOrSkipped[SubmitRunResponse], error) {
	submitRunResponse, err := a.JobsService.Submit(ctx, submitRun)
	if err != nil {
		return nil, err
	}
	return &WaitGetRunJobTerminatedOrSkipped[SubmitRunResponse]{
		Response: submitRunResponse,
		RunId:    submitRunResponse.RunId,
		Poll: func(timeout time.Duration, callback func(*Run)) (*Run, error) {
			return a.WaitGetRunJobTerminatedOrSkipped(ctx, submitRunResponse.RunId, timeout, callback)
		},
		timeout:  20 * time.Minute,
		callback: nil,
	}, nil
}

// Calls [JobsAPI.Submit] and waits to reach TERMINATED or SKIPPED state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[Run](60*time.Minute) functional option.
//
// Deprecated: use [JobsAPI.Submit].Get() or [JobsAPI.WaitGetRunJobTerminatedOrSkipped]
func (a *JobsAPI) SubmitAndWait(ctx context.Context, submitRun SubmitRun, options ...retries.Option[Run]) (*Run, error) {
	wait, err := a.Submit(ctx, submitRun)
	if err != nil {
		return nil, err
	}
	tmp := &retries.Info[Run]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(tmp)
	}
	wait.timeout = tmp.Timeout
	wait.callback = func(info *Run) {
		for _, o := range options {
			o(&retries.Info[Run]{
				Info:    info,
				Timeout: wait.timeout,
			})
		}
	}
	return wait.Get()
}
