{
    "attributes": [
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": null,
            "channel": null,
            "creation_only": null,
            "default_order": null,
            "default_value": null,
            "deprecated": null,
            "description": "Selector is the main selector for the DependencyMapSubview.",
            "exposed": true,
            "filterable": false,
            "foreign_key": null,
            "format": null,
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "selector",
            "orderable": true,
            "primary_key": null,
            "read_only": false,
            "required": null,
            "secret": false,
            "setter": null,
            "stored": true,
            "subtype": "dependencymapview_selector",
            "transient": false,
            "type": "external",
            "unique": null,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": null,
            "channel": null,
            "creation_only": null,
            "default_order": null,
            "default_value": null,
            "deprecated": null,
            "description": "SubSelectors are the selector to apply inside the main selector.",
            "exposed": true,
            "filterable": false,
            "foreign_key": null,
            "format": null,
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "subSelectors",
            "orderable": true,
            "primary_key": null,
            "read_only": false,
            "required": null,
            "secret": false,
            "setter": null,
            "stored": true,
            "subtype": "dependencymapview_subselector",
            "transient": false,
            "type": "external",
            "unique": null,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": null,
            "channel": null,
            "creation_only": null,
            "default_order": null,
            "default_value": null,
            "deprecated": null,
            "description": "Tonality sets the color tonality to use for the DependencyMapSubView.",
            "exposed": true,
            "filterable": true,
            "foreign_key": null,
            "format": "free",
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "tonality",
            "orderable": true,
            "primary_key": null,
            "read_only": false,
            "required": null,
            "secret": false,
            "setter": null,
            "stored": true,
            "subtype": null,
            "transient": false,
            "type": "string",
            "unique": null,
            "uniqueScope": null
        }
    ],
    "children": [],
    "model": {
        "aliases": [],
        "create": null,
        "delete": false,
        "description": "[nodoc]",
        "entity_name": "DependencyMapSubview",
        "extends": [],
        "get": false,
        "package": "Vizualization",
        "resource_name": "dependencymapsubviews",
        "rest_name": "dependencymapsubview",
        "root": null,
        "update": false
    }
}