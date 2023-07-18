package sseserver

import (
	"time"

	"github.com/GeertJohan/go.rice/embedded"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "admin.html",
		FileModTime: time.Unix(1689695249, 0),

		Content: string("<!DOCTYPE html>\n<html lang=\"en\">\n  <head>\n    <title>ssestreamer admin</title>\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1\" />\n\n    <link rel=\"stylesheet\" href=\"//cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/3.1.1/css/bootstrap.min.css\" />\n    <link rel=\"stylesheet\" href=\"//cdnjs.cloudflare.com/ajax/libs/font-awesome/4.1.0/css/font-awesome.min.css\" />\n\n    <script src=\"//cdnjs.cloudflare.com/ajax/libs/jquery/2.1.1/jquery.min.js\"></script>\n    <script src=\"//cdnjs.cloudflare.com/ajax/libs/jquery-timeago/1.4.0/jquery.timeago.min.js\"></script>\n    <script src=\"//cdnjs.cloudflare.com/ajax/libs/knockout/3.1.0/knockout-min.js\"></script>\n    <script src=\"//cdnjs.cloudflare.com/ajax/libs/numeral.js/1.5.3/numeral.min.js\"></script>\n\n    <style media=\"screen\">\n      @media (max-width: 640px) {\n        body{font-size: 12px;}\n      }\n    </style>\n  </head>\n  <body>\n    <article class=\"container\">\n      <h2>system info</h2>\n      <ul>\n        <li>uptime: <time class=\"timeago\" data-bind=\"text: startup_time, attr: {datetime: startup_time}\"></time></li>\n        <li>messages broadcast (unique): <span data-bind=\"text: msgs_broadcast\"></span></li>\n        <!-- <li>memory usage: <span data-bind=\"text: mem_usage\"></span></li>\n        <li>goroutines: <span data-bind=\"text: goroutines\"></span></li> -->\n      </ul>\n      <h2>open streams <small><span data-bind=\"text: clients().length\">N</span> open</small></h2>\n\n      <table class=\"table table-striped table-hover table-condensed\">\n        <thead>\n          <th><i class=\"fa fa-tag\"></i>&nbsp;namespace</th>\n          <th><i class=\"fa fa-globe\"></i>&nbsp;client</th>\n          <th class=\"hidden-xs\"><i class=\"fa fa-flag\"></i>&nbsp;established</th>\n          <th class=\"hidden-xs\"><i class=\"fa fa-check\"></i>&nbsp;msgs</th>\n          <th><i class=\"fa fa-clock-o\"></i>&nbsp;age</th>\n        </thead>\n        <tbody data-bind=\"foreach: clients\">\n          <tr>\n            <td>\n              <code data-bind=\"text: request_path\"></code>\n            </td>\n            <td>\n              <i class=\"{{icon}}\" title=\"{{user_agent}}\"></i>\n              &nbsp;\n              <tt data-bind=\"text: remote_ip\"></tt>\n            </td>\n            <td class=\"hidden-xs\">\n              <time data-bind=\"text: created_at, attr: {datetime: created_iso}\"></time>\n            </td>\n            <td class=\"hidden-xs\">\n              <tt data-bind=\"text: msgs_sent\"></tt>\n            </td>\n            <td>\n              <i class=\"fa fa-clock-o hidden-xs\"></i>\n              <time class=\"timeago\" data-bind=\"text: created_at, attr: {datetime: created_iso}\"></time>\n            </td>\n          </tr>\n        </tbody>\n      </table>\n    </article>\n\n    <script type=\"text/javascript\">\n      // override timeago settings for shorter format\n      jQuery.timeago.settings.strings = {\n        prefixAgo: null,\n        prefixFromNow: null,\n        suffixAgo: \"\",\n        suffixFromNow: \"\",\n        seconds: \"1m\",\n        minute: \"1m\",\n        minutes: \"%dm\",\n        hour: \"1h\",\n        hours: \"%dh\",\n        day: \"1d\",\n        days: \"%dd\",\n        month: \"1mo\",\n        months: \"%dmo\",\n        year: \"1yr\",\n        years: \"%dyr\",\n        wordSeparator: \" \",\n        numbers: []\n      };\n\n      // render status from server\n      var retrieveData = function() {\n        $.get(\"/admin/status.json\", function(response) {\n          adminViewModel.msgs_broadcast( numeral(response.msgs_broadcast).format('0.00a').toUpperCase() );\n          adminViewModel.startup_time( new Date(response.startup_time * 1000).toISOString() );\n\n          response.connections.forEach( function(entry) {\n            adminViewModel.clients.push( new Connection(entry) );\n          });\n\n          // apply timeago (one time only for now, will have to adjust this method\n          // when we want to add new connections on fly)\n          $(\"time.timeago\").timeago();\n\n        }, \"json\");\n      }\n\n      //create KO view model for stats\n      var adminViewModel = {\n        startup_time: ko.observable('-'),\n        msgs_broadcast: ko.observable('-'),\n        mem_usage: ko.observable('-'),\n        goroutines: ko.observable('-'),\n\n        clients: ko.observableArray()\n      }\n\n      function Connection(client) {\n        var self = this;\n        self.request_path = client.namespace;\n        self.remote_ip = client.client_ip;\n        self.user_agent = client.user_agent;\n        self.created_at = new Date(client.created_at * 1000);\n        self.created_iso = new Date(client.created_at * 1000).toISOString();\n        self.msgs_sent = ko.observable(client.msgs_sent);\n      }\n\n      // page is loaded, do stuff\n      $(document).ready( function() {\n        // render the current server status to table\n        retrieveData();\n\n        // apply KO bindings (will cause data to be rendered)\n        ko.applyBindings(adminViewModel);\n      });\n    </script>\n  </body>\n</html>\n"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1689695249, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // "admin.html"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`views`, &embedded.EmbeddedBox{
		Name: `views`,
		Time: time.Unix(1689695249, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"admin.html": file2,
		},
	})
}
