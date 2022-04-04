# iRacing documentation

```json
{
  "car": {
    "assets": {
      "link": "https://members-ng.iracing.com/data/car/assets",
      "note": [
        "image paths are relative to https://images-static.iracing.com/"
      ]
    },
    "get": {
      "link": "https://members-ng.iracing.com/data/car/get"
    }
  },
  "carclass": {
    "get": {
      "link": "https://members-ng.iracing.com/data/carclass/get"
    }
  },
  "constants": {
    "divisions": {
      "link": "https://members-ng.iracing.com/data/constants/divisions",
      "note": "Constant; returned directly as an array of objects"
    }
  },
  "league": {
    "get": {
      "link": "https://members-ng.iracing.com/data/league/get",
      "parameters": {
        "league_id": {
          "type": "number",
          "required": true
        },
        "include_licenses": {
          "type": "boolean",
          "note": "For faster responses, only request when necessary."
        }
      }
    }
  },
  "lookup": {
    "get": {
      "link": "https://members-ng.iracing.com/data/lookup/get",
      "note": "?weather=weather_wind_speed_units&weather=weather_wind_speed_max&weather=weather_wind_speed_min&licenselevels=licenselevels"
    },
    "licenses": {
      "link": "https://members-ng.iracing.com/data/lookup/licenses"
    }
  },
  "member": {
    "get": {
      "link": "https://members-ng.iracing.com/data/member/get",
      "parameters": {
        "cust_ids": {
          "type": "numbers",
          "required": true,
          "note": "?cust_ids=2,3,4"
        },
        "include_licenses": {
          "type": "boolean"
        }
      }
    },
    "info": {
      "link": "https://members-ng.iracing.com/data/member/info",
      "note": "Always the authenticated member."
    }
  },
  "results": {
    "get": {
      "link": "https://members-ng.iracing.com/data/results/get",
      "note": "Get the results of a subsession, if authorized to view them. series_logo image paths are relative to https://images-static.iracing.com/img/logos/series/",
      "parameters": {
        "subsession_id": {
          "type": "number",
          "required": true
        },
        "include_licenses": {
          "type": "boolean"
        }
      }
    },
    "event_log": {
      "link": "https://members-ng.iracing.com/data/results/event_log",
      "parameters": {
        "subsession_id": {
          "type": "number",
          "required": true
        },
        "simsession_number": {
          "type": "number",
          "required": true,
          "note": "The main event is 0; the preceding event is -1, and so on."
        }
      }
    },
    "lap_chart_data": {
      "link": "https://members-ng.iracing.com/data/results/lap_chart_data",
      "parameters": {
        "subsession_id": {
          "type": "number",
          "required": true
        },
        "simsession_number": {
          "type": "number",
          "required": true,
          "note": "The main event is 0; the preceding event is -1, and so on."
        }
      }
    },
    "lap_data": {
      "link": "https://members-ng.iracing.com/data/results/lap_data",
      "parameters": {
        "subsession_id": {
          "type": "number",
          "required": true
        },
        "simsession_number": {
          "type": "number",
          "required": true,
          "note": "The main event is 0; the preceding event is -1, and so on."
        },
        "cust_id": {
          "type": "number",
          "note": "Required if the subsession was a single-driver event. Optional for team events. If omitted for a team event then the laps driven by all the team's drivers will be included."
        },
        "team_id": {
          "type": "number",
          "note": "Required if the subsession was a team event."
        }
      }
    },
    "season_results": {
      "link": "https://members-ng.iracing.com/data/results/season_results",
      "parameters": {
        "season_id": {
          "type": "number",
          "required": true
        },
        "event_type": {
          "type": "number",
          "note": "Retrict to one event type: 2 - Practice; 3 - Qualify; 4 - Time Trial; 5 - Race"
        },
        "race_week_num": {
          "type": "number",
          "note": "The first race week of a season is 0."
        }
      }
    }
  },
  "series": {
    "seasons": {
      "link": "https://members-ng.iracing.com/data/series/seasons",
      "parameters": {
        "include_series": {
          "type": "boolean"
        }
      }
    },
    "stats_series": {
      "link": "https://members-ng.iracing.com/data/series/stats_series",
      "note": "To get series and seasons for which standings should be available, filter the list by official: true."
    }
  },
  "stats": {
    "member_career": {
      "link": "https://members-ng.iracing.com/data/stats/member_career",
      "parameters": {
        "cust_id": {
          "type": "number",
          "note": "Defaults to the authenticated member."
        }
      }
    },
    "member_division": {
      "link": "https://members-ng.iracing.com/data/stats/member_division",
      "note": "Divisions are 0-based: 0 is Division 1, 10 is Rookie. See /data/constants/divisons for more information. Always for the authenticated member.",
      "parameters": {
        "season_id": {
          "type": "number",
          "required": true
        },
        "event_type": {
          "type": "number",
          "required": true,
          "note": "The event type code for the division type: 4 - Time Trial; 5 - Race"
        }
      }
    },
    "member_recent_races": {
      "link": "https://members-ng.iracing.com/data/stats/member_recent_races",
      "parameters": {
        "cust_id": {
          "type": "number",
          "note": "Defaults to the authenticated member."
        }
      }
    },
    "member_summary": {
      "link": "https://members-ng.iracing.com/data/stats/member_summary",
      "parameters": {
        "cust_id": {
          "type": "number",
          "note": "Defaults to the authenticated member."
        }
      }
    },
    "member_yearly": {
      "link": "https://members-ng.iracing.com/data/stats/member_yearly",
      "parameters": {
        "cust_id": {
          "type": "number",
          "note": "Defaults to the authenticated member."
        }
      }
    },
    "season_driver_standings": {
      "link": "https://members-ng.iracing.com/data/stats/season_driver_standings",
      "parameters": {
        "season_id": {
          "type": "number",
          "required": true
        },
        "car_class_id": {
          "type": "number",
          "required": true
        },
        "race_week_num": {
          "type": "number",
          "note": "The first race week of a season is 0."
        }
      }
    },
    "season_supersession_standings": {
      "link": "https://members-ng.iracing.com/data/stats/season_supersession_standings",
      "parameters": {
        "season_id": {
          "type": "number",
          "required": true
        },
        "car_class_id": {
          "type": "number",
          "required": true
        },
        "race_week_num": {
          "type": "number",
          "note": "The first race week of a season is 0."
        }
      }
    },
    "season_team_standings": {
      "link": "https://members-ng.iracing.com/data/stats/season_team_standings",
      "parameters": {
        "season_id": {
          "type": "number",
          "required": true
        },
        "car_class_id": {
          "type": "number",
          "required": true
        },
        "race_week_num": {
          "type": "number",
          "note": "The first race week of a season is 0."
        }
      }
    },
    "season_tt_standings": {
      "link": "https://members-ng.iracing.com/data/stats/season_tt_standings",
      "parameters": {
        "season_id": {
          "type": "number",
          "required": true
        },
        "car_class_id": {
          "type": "number",
          "required": true
        },
        "race_week_num": {
          "type": "number",
          "note": "The first race week of a season is 0."
        }
      }
    },
    "season_tt_results": {
      "link": "https://members-ng.iracing.com/data/stats/season_tt_results",
      "parameters": {
        "season_id": {
          "type": "number",
          "required": true
        },
        "car_class_id": {
          "type": "number",
          "required": true
        },
        "race_week_num": {
          "type": "number",
          "required": true,
          "note": "The first race week of a season is 0."
        }
      }
    },
    "season_qualify_results": {
      "link": "https://members-ng.iracing.com/data/stats/season_qualify_results",
      "parameters": {
        "season_id": {
          "type": "number",
          "required": true
        },
        "car_class_id": {
          "type": "number",
          "required": true
        },
        "race_week_num": {
          "type": "number",
          "required": true,
          "note": "The first race week of a season is 0."
        }
      }
    }
  },
  "track": {
    "assets": {
      "link": "https://members-ng.iracing.com/data/track/assets",
      "note": [
        "image paths are relative to https://images-static.iracing.com/"
      ]
    },
    "get": {
      "link": "https://members-ng.iracing.com/data/track/get"
    }
  }
}
```
