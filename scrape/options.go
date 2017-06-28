package scrape

import (
	"time"
)

// ScrapeOptions contains options that are used during the progress of a
// scrape.
type ScrapeOptions struct {
	// The maximum number of pages to scrape.  The scrape will proceed until
	// either this number of pages have been scraped, or until the paginator
	// returns no further URLs.  Set this value to 0 to indicate an unlimited
	// number of pages can be scraped.
	MaxPages int
	//Output format 
	Format string
	//CrawlDelay should be used for a scraper to throttle the crawling speed to avoid hitting the web servers too frequently. 
	FetchDelay time.Duration
	//Some sites track  statistically significant similarities in the time between requests to them. RandomizeCrawlDelay setting decreases the chance of a crawler being blocked by such sites. This way a random delay ranging from 0.5 * CrawlDelay to 1.5 * CrawlDelay seconds is used between consecutive requests to the same domain. If CrawlDelay is zero (default) this option has no effect.
	RandomizeFetchDelay bool
	//paginated results are returned. Single list of combined results from every block on all pages is returned by default. Paginated results is actual for JSON and XML formats. Combined list of results is always returned for CSV format.  
	PaginatedResults bool
}	

// The default options during a scrape.
var DefaultOptions = ScrapeOptions{
	MaxPages: 0,
	Format: "json",
	PaginatedResults: false,
	FetchDelay: 1,
	RandomizeFetchDelay: true,
}