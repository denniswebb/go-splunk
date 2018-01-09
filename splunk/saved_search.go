package splunk

import (
	"fmt"
)

type SavedSearch struct {
	Name   string `schema:"name" xml:"name"`
	Search string `schema:"search" xml:"search"`

	// Read-only attribute. The state of the email action. Value ignored on POST. Use Actions to specify a list of enabled actions.
	ActionEmail bool `schema:"action_email" xml:"action.email"`

	// The password to use when authenticating with the SMTP server. Normally this value will be set when editing the email settings, however you can set a clear text password here and it will be encrypted on the next Splunk restart.Defaults to empty string.
	ActionEmailAuthPassword string `schema:"action_email_auth_password" xml:"action.email.auth_password"`

	// BCC email address to use ction.email.cc" xml:"action.email.cc"`
	ActionEmailBCC string `schema:"action_email_bcc" xml:"action.email.bcc"`

	// CC email address to use if action.email is enabled.
	ActionEmailCC string `schema:"action_email_cc" xml:"action.email.cc"`

	// The username to use when authenticating with the SMTP server.
	ActionEmailAuthUsername string `schema:"action_email_auth_username" xml:"action.email.auth_username"`

	// The search command (or pipeline) which is responsible for executing the action.Generally the command is a template search pipeline which is realized with values from the saved search. To reference saved search field values wrap them in $, for example to reference the savedsearch name use $name$, to reference the search use $search$.
	ActionEmailCommand string `schema:"action_email_command" xml:"action.email.command"`

	// Valid values: (plain | html | raw | csv)Specify the format of text in the email. This value also applies to any attachments.
	ActionEmailFormat string `schema:"action_email_format" xml:"action.email.format"`

	// Email address from which the email action originates.Defaults to splunk@$LOCALHOST or whatever value is set in alert_actions.conf.
	ActionEmailFrom string `schema:"action_email_from" xml:"action.email.from"`

	// Sets the hostname used in the web link (url) sent in email actions.
	ActionEmailHostname string `schema:"action_email_hostname" xml:"action.email.hostname"`

	// Indicates whether the search results are contained in the body of the email.Results can be either inline or attached to an email. See action.email.sendresults.
	ActionEmailInline bool `schema:"action_email_inline" xml:"action.email.inline"`

	// Set the address of the MTA server to be used to send the emails.Defaults to <LOCALHOST> (or whatever is set in alert_actions.conf).
	ActionEmailMailserver string `schema:"action_email_mailserver" xml:"action.email.mailserver"`

	// Sets the global maximum number of search results to send when email.action is enabled.
	ActionEmailMaxResults int `schema:"action_email_maxresults" xml:"action.email.maxresults"`

	// Valid values are Integer[m|s|h|d].Specifies the maximum amount of time the execution of an email action takes before the action is aborted. Defaults to 5m.
	ActionEmailMaxTime string `schema:"action_email_maxtime" xml:"action.email.maxtime"`

	// The name of the view to deliver if sendpdf is enabled
	ActionEmailPDFView string `schema:"action_email_pdfview" xml:"action.email.pdfview"`

	// Search string to preprocess results before emailing them. Defaults to empty string (no preprocessing).Usually the preprocessing consists of filtering out unwanted internal fields.
	ActionEmailPreprocessResults string `schema:"action_email_preprocess_results" xml:"action.email.preprocess_results"`

	// Space-separated list.
	ActionEmailReportCIDFontList string `schema:"action_email_reportCIDFontList" xml:"action.email.reportCIDFontList"`

	// Indicates whether to include the Splunk logo with the report.
	ActionEmailReportIncludeSplunkLogo bool `schema:"action_email_reportIncludeSplunkLogo" xml:"action.email.reportIncludeSplunkLogo"`

	// Valid values: (portrait | landscape)Specifies the paper orientation: portrait or landscape. Defaults to portrait.
	ActionEmailReportPaperOrientation string `schema:"action_email_reportPaperOrientation" xml:"action.email.reportPaperOrientation"`

	// Valid values: (letter | legal | ledger | a2 | a3 | a4 | a5)Specifies the paper size for PDFs. Defaults to letter.
	ActionEmailReportPaperSize string `schema:"action_email_reportPaperSize" xml:"action.email.reportPaperSize"`

	// Indicates whether the PDF server is enabled. Defaults to false.
	ActionEmailReportServerEnabled bool `schema:"action_email_reportServerEnabled" xml:"action.email.reportServerEnabled"`

	// The URL of the PDF report server, if one is set up and available on the network.For a default locally installed report server, the URL is http: localhost:8091/
	ActionEmailReportServerURL string `schema:"action_email_reportServerURL" xml:"action.email.reportServerURL"`

	// Indicates whether to create and send the results as a PDF. Defaults to false.
	ActionEmailSendPDF bool `schema:"action_email_sendpdf" xml:"action.email.sendpdf"`

	// Indicates whether to attach the search results in the email.Results can be either attached or inline. See action.email.inline.
	ActionEmailSendResults bool `schema:"action_email_sendresults" xml:"action.email.sendresults"`

	// Specifies an alternate email subject.Defaults to SplunkAlert-<savedsearchname>.
	ActionEmailSubject string `schema:"action_email_subject" xml:"action.email.subject"`

	// A comma or semicolon separated list of recipient email addresses. Required if this search is scheduled and the email alert action is enabled.
	ActionEmailTo string `schema:"action_email_to" xml:"action.email.to"`

	// Indicates whether the execution of this action signifies a trackable alert.
	ActionEmailTrackAlert bool `schema:"action_email_track_alert" xml:"action.email.track_alert"`

	// Specifies the minimum time-to-live in seconds of the search artifacts if this action is triggered.
	ActionEmailTTL string `schema:"action_email_ttl" xml:"action.email.ttl"`

	// Indicates whether to use SSL when communicating with the SMTP server.Defaults to false.
	ActionEmailUseSSL bool `schema:"action_email_use_ssl" xml:"action.email.use_ssl"`

	// Indicates whether to use TLS (transport layer security) when communicating with the SMTP server (starttls).Defaults to false.
	ActionEmailUseTLS bool `schema:"action_email_use_tls" xml:"action.email.use_tls"`

	// Indicates whether columns should be sorted from least wide to most wide, left to right.Only valid if format=text.
	ActionEmailWidthSortColumns bool `schema:"action_email_width_sort_columns" xml:"action.email.width_sort_columns"`

	// Read-only attribute. The state of the populate lookup action. Value ignored on POST. Use actions to specify a list of enabled actions.
	ActionPopulateLookup bool `schema:"action_populate_lookup" xml:"action.populate_lookup"`

	// The search command (or pipeline) which is responsible for executing the action.Generally the command is a template search pipeline which is realized with values from the saved search. To reference saved search field values wrap them in $, for example to reference the savedsearch name use $name$, to reference the search use $search$.
	ActionPopulateLookupCommand string `schema:"action_populate_lookup_command" xml:"action.populate_lookup.command"`

	// Lookup name of path of the lookup to populate
	ActionPopulateLookupDest string `schema:"action_populate_lookup_dest" xml:"action.populate_lookup.dest"`

	// Sets the hostname used in the web link (url) sent in alert actions.
	ActionPopulateLookupHostname string `schema:"action_populate_lookup_hostname" xml:"action.populate_lookup.hostname"`

	// Sets the maximum number of search results sent using alerts. Defaults to 100.
	ActionPopulateLookupMaxResults int `schema:"action_populate_lookup_maxresults" xml:"action.populate_lookup.maxresults"`

	// Valid values are: Integer[m|s|h|d]Sets the maximum amount of time the execution of an action takes before the action is aborted. Defaults to 5m.
	ActionPopulateLookupMaxTime string `schema:"action_populate_lookup_maxtime" xml:"action.populate_lookup.maxtime"`

	// Indicates whether the execution of this action signifies a trackable alert.
	ActionPopulateLookupTrackAlert bool `schema:"action_populate_lookup_track_alert" xml:"action.populate_lookup.track_alert"`

	// Specifies the minimum time-to-live in seconds of the search artifacts if this action is triggered.
	ActionPopulateLookupTTL string `schema:"action_populate_lookup_ttl" xml:"action.populate_lookup.ttl"`

	// Read-only attribute. The state of the rss action. Value ignored on POST. Use actions to specify a list of enabled actions.
	ActionRSS bool `schema:"action_rss" xml:"action.rss"`

	// The search command (or pipeline) which is responsible for executing the action.Generally the command is a template search pipeline which is realized with values from the saved search. To reference saved search field values wrap them in $, for example to reference the savedsearch name use $name$, to reference the search use $search$.
	ActionRSSCommand string `schema:"action_rss_command" xml:"action.rss.command"`

	// Sets the hostname used in the web link (url) sent in alert actions.
	ActionRSSHostname string `schema:"action_rss_hostname" xml:"action.rss.hostname"`

	// Sets the maximum number of search results sent using alerts. Defaults to 100.
	ActionRSSMaxResults int `schema:"action_rss_maxresults" xml:"action.rss.maxresults"`

	// Valid values are Integer[m|s|h|d].Sets the maximum amount of time the execution of an action takes before the action is aborted. Defaults to 1m.
	ActionRSSMaxTime string `schema:"action_rss_maxtime" xml:"action.rss.maxtime"`

	// Indicates whether the execution of this action signifies a trackable alert.
	ActionRSSTrackAlert bool `schema:"action_rss_track_alert" xml:"action.rss.track_alert"`

	// Specifies the minimum time-to-live in seconds of the search artifacts if this action is triggered.
	ActionRSSTTL string `schema:"action_rss_ttl" xml:"action.rss.ttl"`

	// Read-only attribute. The state of the script action. Value ignored on POST. Use actions to specify a list of enabled actions.
	ActionScript bool `schema:"action_script" xml:"action.script"`

	// The search command (or pipeline) which is responsible for executing the action.Generally the command is a template search pipeline which is realized with values from the saved search. To reference saved search field values wrap them in $, for example to reference the savedsearch name use $name$, to reference the search use $search$.
	ActionScriptCommand string `schema:"action_script_command" xml:"action.script.command"`

	// File name of the script to call. Required if script action is enabled
	ActionScriptFilename string `schema:"action_script_filename" xml:"action.script.filename"`

	// Sets the hostname used in the web link (url) sent in alert actions.
	ActionScriptHostname string `schema:"action_script_hostname" xml:"action.script.hostname"`

	// Sets the maximum number of search results sent using alerts. Defaults to 100.
	ActionScriptMaxResults int `schema:"action_script_maxresults" xml:"action.script.maxresults"`

	// Valid values are: Integer[m|s|h|d]Sets the maximum amount of time the execution of an action takes before the action is aborted. Defaults to 5m.
	ActionScriptMaxTime string `schema:"action_script_maxtime" xml:"action.script.maxtime"`

	// Indicates whether the execution of this action signifies a trackable alert.
	ActionScriptTrackAlert bool `schema:"action_script_track_alert" xml:"action.script.track_alert"`

	// Specifies the minimum time-to-live in seconds of the search artifacts if this action is triggered.
	ActionScriptTTL string `schema:"action_script_ttl" xml:"action.script.ttl"`

	// Read-only attribute. The state of the summary index action. Value ignored on POST. Use actions to specify a list of enabled actions.
	ActionSummaryIndex int `schema:"action_summary_index" xml:"action.summary_index"`

	// Specifies the name of the summary index where the results of the scheduled search are saved.Defaults to "summary."
	ActionSummaryIndexName string `schema:"action_summary_index__name" xml:"action.summary_index._name"`

	// The search command (or pipeline) which is responsible for executing the action.
	ActionSummaryIndexCommand string `schema:"action_summary_index_command" xml:"action.summary_index.command"`

	// Sets the hostname used in the web link (url) sent in summary-index alert actions.
	ActionSummaryIndexHostname string `schema:"action_summary_index_hostname" xml:"action.summary_index.hostname"`

	// Determines whether to execute the summary indexing action as part of the scheduled search.
	ActionSummaryIndexInline bool `schema:"action_summary_index_inline" xml:"action.summary_index.inline"`

	// Sets the maximum number of search results sent using alerts. Defaults to 100.
	ActionSummaryIndexMaxResults int `schema:"action_summary_index_maxresults" xml:"action.summary_index.maxresults"`

	// Valid values are: Integer[m|s|h|d]Sets the maximum amount of time the execution of an action takes before the action is aborted. Defaults to 5m.
	ActionSummaryIndexMaxTime string `schema:"action_summary_index_maxtime" xml:"action.summary_index.maxtime"`

	// Indicates whether the execution of this action signifies a trackable alert.
	ActionSummaryIndexTrackAlert bool `schema:"action_summary_index_track_alert" xml:"action.summary_index.track_alert"`

	// Specifies the minimum time-to-live in seconds of the search artifacts if this action is triggered.
	ActionSummaryIndexTTL string `schema:"action_summary_index_ttl" xml:"action.summary_index.ttl"`

	// A comma-separated list of actions to enable.For example: rss,email
	Actions string `schema:"actions" xml:"actions"`

	// Specifies whether Splunk applies the alert actions to the entire result set or on each individual result. Defaults to true.
	AlertDigestMode bool `schema:"alert_digest_mode" xml:"alert.digest_mode"`

	// Sets the period of time to show the alert in the dashboard. For example: 60 = 60 seconds, 1m = 1 minute, 1h = 60 minutes = 1 hour.
	AlertExpires string `schema:"alert_expires" xml:"alert.expires"`

	// Sets the alert severity level. Valid values: (1 | 2 | 3 | 4 | 5 | 6). 1 DEBUG 2 INFO 3 WARN 4 ERROR 5 SEVERE 6 FATAL
	AlertSeverity int `schema:"alert_severity" xml:"alert.severity"`

	// Indicates whether alert suppression is enabled for this scheduled search.
	AlertSuppress bool `schema:"alert_suppress" xml:"alert.suppress"`

	// Comma delimited list of fields to use for suppression when doing per result alerting. Required if suppression is turned on and per result alerting is enabled.
	AlertSuppressFields string `schema:"alert_suppress_fields" xml:"alert.suppress.fields"`

	// Specifies the suppresion period. Only valid if alert.supress is enabled. For example: 60 = 60 seconds, 1m = 1 minute, 1h = 60 minutes = 1 hour.
	AlertSuppressPeriod string `schema:"alert_suppress_period" xml:"alert.suppress.period"`

	// Specifies whether to track the actions triggered by this scheduled search. Valid values: (true | false | auto)
	AlertTrack string `schema:"alert_track" xml:"alert.track"`

	// One of the following strings: greater than, less than, equal to, rises by, drops by, rises by perc, drops by percUsed with alert_threshold to trigger alert actions.
	AlertComparator string `schema:"alert_comparator" xml:"alert_comparator"`

	// Contains a conditional search that is evaluated against the results of the saved search.
	AlertCondition string `schema:"alert_condition" xml:"alert_condition"`

	// Valid values are: Integer[%]Specifies the value to compare (see alert_comparator) before triggering the alert actions. If expressed as a percentage, indicates value to use when alert_comparator is set to "rises by perc" or "drops by perc."
	AlertThreshold string `schema:"alert_threshold" xml:"alert_threshold"`

	// What to base the alert on, overriden by alert_condition if it is specified. Valid values are: always, custom, number of events, number of hosts, number of sources
	AlertType string `schema:"alert_type" xml:"alert_type"`

	// Indicates whether the scheduler should ensure that the data for this search is automatically summarized.
	AutoSummarize bool `schema:"auto_summarize" xml:"auto_summarize"`

	// A search template that constructs the auto summarization for this search.Caution: Advanced feature. Do not change unless you understand the architecture of auto summarization of saved searches.
	AutoSummarizeCommand string `schema:"auto_summarize_command" xml:"auto_summarize.command"`

	// Cron schedule that probes and generates the summaries for this saved search.The default value corresponds to every ten minutes.
	AutoSummarizeCronSchedule string `schema:"auto_summarize_cron_schedule" xml:"auto_summarize.cron_schedule"`

	// A time string that specifies the earliest time for summarizing this search. Can be a relative or absolute time.If this value is an absolute time, use the dispatch.time_format to format the value.
	AutoSummarizeDispatchEarliestTime string `schema:"auto_summarize_dispatch_earliest_time" xml:"auto_summarize.dispatch.earliest_time"`

	// A time string that specifies the latest time for summarizing this saved search. Can be a relative or absolute time.If this value is an absolute time, use the dispatch.time_format to format the value.
	AutoSummarizeDispatchLatestTime string `schema:"auto_summarize_dispatch_latest_time" xml:"auto_summarize.dispatch.latest_time"`

	// Defines the time format that Splunk uses to specify the earliest and latest time.
	AutoSummarizeDispatchTimeFormat string `schema:"auto_summarize_dispatch_time_format" xml:"auto_summarize.dispatch.time_format"`

	// Valid values: Integer[p]Indicates the time to live (in seconds) for the artifacts of the summarization of the scheduled search.
	AutoSummarizeDispatchTTL string `schema:"auto_summarize_dispatch_ttl" xml:"auto_summarize.dispatch.ttl"`

	// The maximum number of buckets with the suspended summarization before the summarization search is completely stopped, and the summarization of the search is suspended for auto_summarize.suspend_period.
	AutoSummarizeMaxDisabledBuckets int `schema:"auto_summarize_max_disabled_buckets" xml:"auto_summarize.max_disabled_buckets"`

	// The maximum ratio of summary_size/bucket_size, which specifies when to stop summarization and deem it unhelpful for a bucket.Note: The test is only performed if the summary size is larger than auto_summarize.max_summary_size.
	AutoSummarizeMaxSummaryRatio float64 `schema:"auto_summarize_max_summary_ratio" xml:"auto_summarize.max_summary_ratio"`

	// The minimum summary size, in bytes, before testing whether the summarization is helpful.The default value is equivalent to 5MB.
	AutoSummarizeMaxSummarySize int `schema:"auto_summarize_max_summary_size" xml:"auto_summarize.max_summary_size"`

	// Maximum time (in seconds) that the summary search is allowed to run.Note: This is an approximate time. The summary search stops at clean bucket boundaries.
	AutoSummarizeMaxTime int `schema:"auto_summarize_max_time" xml:"auto_summarize.max_time"`

	// Time specfier indicating when to suspend summarization of this search if the summarization is deemed unhelpful.
	AutoSummarizeSuspendPeriod string `schema:"auto_summarize_suspend_period" xml:"auto_summarize.suspend_period"`

	// The list of time ranges that each summarized chunk should span. This comprises the list of available granularity levels for which summaries would be available. Specify a comma delimited list of time specifiers.For example a timechart over the last month whose granuality is at the day level should set this to 1d. If you need the same data summarized at the hour level for weekly charts, use: 1h,1d.
	AutoSummarizeTimespan string `schema:"auto_summarize_timespan" xml:"auto_summarize.timespan"`

	// cron stringThe cron schedule to execute this search.
	CronSchedule string `schema:"cron_schedule" xml:"cron_schedule"`

	// Human-readable description of this saved search. Defaults to empty string.
	Description string `schema:"description" xml:"description"`

	// Indicates if the saved search is enabled.Disabled saved searches are not visible in Splunk Web.
	Disabled bool `schema:"disabled" xml:"disabled"`

	// The maximum number of timeline buckets.
	DispatchBuckets int `schema:"dispatch_buckets" xml:"dispatch.buckets"`

	// A time string that specifies the earliest time for this search. Can be a relative or absolute time.If this value is an absolute time, use the dispatch.time_format to format the value.
	DispatchEarliestTime string `schema:"dispatch_earliest_time" xml:"dispatch.earliest_time"`

	// Indicates whether to used indexed-realtime mode when doing real-time searches.
	DispatchIndexedRealtime bool `schema:"dispatch_indexedRealtime" xml:"dispatch.indexedRealtime"`

	// A time string that specifies the latest time for this saved search. Can be a relative or absolute time.If this value is an absolute time, use the dispatch.time_format to format the value.
	DispatchLatestTime string `schema:"dispatch_latest_time" xml:"dispatch.latest_time"`

	// Enables or disables the lookups for this search.
	DispatchLookups bool `schema:"dispatch_lookups" xml:"dispatch.lookups"`

	// The maximum number of results before finalizing the search.
	DispatchMaxCount int `schema:"dispatch_max_count" xml:"dispatch.max_count"`

	// Indicates the maximum amount of time (in seconds) before finalizing the search.
	DispatchMaxTime int `schema:"dispatch_max_time" xml:"dispatch.max_time"`

	// Specifies, in seconds, how frequently Splunk should run the MapReduce reduce phase on accumulated map values.
	DispatchReduceFreq int `schema:"dispatch_reduce_freq" xml:"dispatch.reduce_freq"`

	// Whether to back fill the real time window for this search. Parameter valid only if this is a real time search
	DispatchRtBackfill bool `schema:"dispatch_rt_backfill" xml:"dispatch.rt_backfill"`

	// Specifies whether Splunk spawns a new search process when this saved search is executed.Searches against indexes must run in a separate process.
	DispatchSpawnProcess bool `schema:"dispatch_spawn_process" xml:"dispatch.spawn_process"`

	// A time format string that defines the time format that Splunk uses to specify the earliest and latest time.
	DispatchTimeFormat string `schema:"dispatch_time_format" xml:"dispatch.time_format"`

	// Indicates the time to live (in seconds) for the artifacts of the scheduled search, if no actions are triggered.
	DispatchTTL string `schema:"dispatch_ttl" xml:"dispatch.ttl"`

	// Defines the default UI view name (not label) in which to load the results. Accessibility is subject to the user having sufficient permissions.
	DisplayView string `schema:"displayview" xml:"displayview"`

	// Whether this search is to be run on a schedule
	IsScheduled bool `schema:"is_scheduled" xml:"is_scheduled"`

	// Specifies whether this saved search should be listed in the visible saved search list.
	IsVisible bool `schema:"is_visible" xml:"is_visible"`

	// The maximum number of concurrent instances of this search the scheduler is allowed to run.
	MaxConcurrent int `schema:"max_concurrent" xml:"max_concurrent"`

	// Read-only attribute. Value ignored on POST. There are some old clients who still send this value
	NextScheduledTime string `schema:"next_scheduled_time" xml:"next_scheduled_time"`

	// Read-only attribute. Value ignored on POST. Splunk computes this value during runtime.
	QualifiedSearch string `schema:"qualifiedSearch" xml:"qualifiedSearch"`

	// Controls the way the scheduler computes the next execution time of a scheduled search.
	RealtimeSchedule bool `schema:"realtime_schedule" xml:"realtime_schedule"`

	// Specifies a field used by Splunk UI to denote the app this search should be dispatched in.
	RequestUIDispatchApp string `schema:"request_ui_dispatch_app" xml:"request.ui_dispatch_app"`

	// Specifies a field used by Splunk UI to denote the view this search should be displayed in.
	RequestUIDispatchView string `schema:"request_ui_dispatch_view" xml:"request.ui_dispatch_view"`

	// Specifies whether to restart a real-time search managed by the scheduler when a search peer becomes available for this saved search.
	RestartOnSearchPeerAdd bool `schema:"restart_on_searchpeer_add" xml:"restart_on_searchpeer_add"`

	// Indicates whether this search runs when Splunk starts. If it does not run on startup, it runs at the next scheduled time.Splunk recommends that you set run_on_startup to true for scheduled searches that populate lookup tables.
	RunOnStartup bool `schema:"run_on_startup" xml:"run_on_startup"`

	// Defines the viewstate id associated with the UI view listed in 'displayview'.
	VSID string `schema:"vsid" xml:"vsid"`
}

func (c *Client) SavedSearchPost(savedSearch *SavedSearch, path string) (r SavedSearch, e error) {
	params, e := encode(savedSearch)
	if e != nil {
		return
	}

	f, e := c.Post(path, params)
	if e != nil {
		return
	}

	r = feedToSavedSearch(f)
	return
}

func (c *Client) SavedSearchCreate(s *SavedSearch) (r SavedSearch, e error) {
	return c.SavedSearchPost(s, PathSavedSearchCreate)
}

func (c *Client) SavedSearchRead(name string) (r SavedSearch, e error) {
	f, e := c.Get(fmt.Sprintf(PathSavedSearchUpdate, name))
	if e != nil {
		return
	}

	r = feedToSavedSearch(f)
	return
}

// SavedSearchDelete deletes a Saved Search from Splunk
func (c *Client) SavedSearchDelete(name string) (e error) {
	return c.Delete(fmt.Sprintf(PathSavedSearchUpdate, name))
}

func (c *Client) SavedSearchUpdate(savedSearch *SavedSearch) (r SavedSearch, e error) {
	params, e := encode(savedSearch)
	if e != nil {
		return
	}

	// Delete name param as it's not allow in Post on Update
	delete(params, "name")
	f, e := c.Post(fmt.Sprintf(PathSavedSearchUpdate, savedSearch.Name), params)
	if e != nil {
		return
	}

	r = feedToSavedSearch(f)
	return
}

func feedToSavedSearch(f Feed) (s SavedSearch) {
	if len(f.Entry) == 0 {
		return SavedSearch{}
	}

	e := f.Entry[0]

	s.Name = e.Title

	decoder.Decode(&s, e.PropertyMap())
	return
}
