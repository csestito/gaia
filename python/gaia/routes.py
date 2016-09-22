# -*- coding: utf-8 -*-

routes = []

# routes for mynamespaces

# routes for mapnodes

# routes for layers
routes.append(("GET", "/layers/:id"))

# routes for dependencymapsubviews

# routes for apiauthorizationpolicies
routes.append(("GET", "/apiauthorizationpolicies/:id"))
routes.append(("PUT", "/apiauthorizationpolicies/:id"))
routes.append(("DELETE", "/apiauthorizationpolicies/:id"))

# routes for images
routes.append(("GET", "/images/:id"))

# routes for systemcalls
routes.append(("GET", "/systemcalls/:id"))
routes.append(("PUT", "/systemcalls/:id"))
routes.append(("DELETE", "/systemcalls/:id"))

# routes for tags

# routes for mapedges

# routes for certificates
routes.append(("POST", "/certificates/:id"))
routes.append(("GET", "/certificates/:id"))
routes.append(("PUT", "/certificates/:id"))
routes.append(("DELETE", "/certificates/:id"))

# routes for filepaths
routes.append(("GET", "/filepaths/:id"))
routes.append(("PUT", "/filepaths/:id"))
routes.append(("DELETE", "/filepaths/:id"))

# routes for namespaces
routes.append(("GET", "/namespaces/:id"))
routes.append(("PUT", "/namespaces/:id"))
routes.append(("DELETE", "/namespaces/:id"))

# routes for policyrules

# routes for externalservices
routes.append(("GET", "/externalservices/:id"))
routes.append(("PUT", "/externalservices/:id"))
routes.append(("DELETE", "/externalservices/:id"))

# routes for policies
routes.append(("GET", "/policies/:id"))
routes.append(("PUT", "/policies/:id"))
routes.append(("DELETE", "/policies/:id"))

# routes for flowstatistics

# routes for fileaccesspolicies
routes.append(("GET", "/fileaccesspolicies/:id"))
routes.append(("PUT", "/fileaccesspolicies/:id"))
routes.append(("DELETE", "/fileaccesspolicies/:id"))

# routes for healthreports

# routes for users
routes.append(("GET", "/users/:id"))
routes.append(("PUT", "/users/:id"))
routes.append(("DELETE", "/users/:id"))
routes.append(("POST", "/users/:id/certificates"))
routes.append(("GET", "/users/:id/certificates"))

# routes for renderedpolicies

# routes for namespacemappingpolicies
routes.append(("GET", "/namespacemappingpolicies/:id"))
routes.append(("PUT", "/namespacemappingpolicies/:id"))
routes.append(("DELETE", "/namespacemappingpolicies/:id"))

# routes for processingunits
routes.append(("GET", "/processingunits/:id"))
routes.append(("PUT", "/processingunits/:id"))
routes.append(("DELETE", "/processingunits/:id"))
routes.append(("GET", "/processingunits/:id/renderedpolicies"))

# routes for authenticators
routes.append(("GET", "/authenticators/:id"))
routes.append(("PUT", "/authenticators/:id"))
routes.append(("DELETE", "/authenticators/:id"))
routes.append(("POST", "/authenticators/:id/users"))
routes.append(("GET", "/authenticators/:id/users"))

# routes for dependencymapviews
routes.append(("GET", "/dependencymapviews/:id"))
routes.append(("PUT", "/dependencymapviews/:id"))
routes.append(("DELETE", "/dependencymapviews/:id"))

# routes for dependencymaps

# routes for vulnerabilities

# routes for rendereddependencymapviews

# routes for servers
routes.append(("GET", "/servers/:id"))
routes.append(("PUT", "/servers/:id"))
routes.append(("DELETE", "/servers/:id"))
routes.append(("POST", "/servers/:id/certificates"))
routes.append(("GET", "/servers/:id/certificates"))

# routes for root
routes.append(("GET", "/root"))
routes.append(("POST", "/apiauthorizationpolicies"))
routes.append(("GET", "/apiauthorizationpolicies"))
routes.append(("POST", "/authenticators"))
routes.append(("GET", "/authenticators"))
routes.append(("GET", "/certificates"))
routes.append(("GET", "/dependencymaps"))
routes.append(("POST", "/dependencymapviews"))
routes.append(("GET", "/dependencymapviews"))
routes.append(("POST", "/externalservices"))
routes.append(("GET", "/externalservices"))
routes.append(("POST", "/fileaccesspolicies"))
routes.append(("GET", "/fileaccesspolicies"))
routes.append(("POST", "/filepaths"))
routes.append(("GET", "/filepaths"))
routes.append(("GET", "/flowstatistics"))
routes.append(("GET", "/images"))
routes.append(("GET", "/layers"))
routes.append(("GET", "/mynamespaces"))
routes.append(("POST", "/namespaces"))
routes.append(("GET", "/namespaces"))
routes.append(("POST", "/namespacemappingpolicies"))
routes.append(("GET", "/namespacemappingpolicies"))
routes.append(("POST", "/networkaccesspolicies"))
routes.append(("GET", "/networkaccesspolicies"))
routes.append(("POST", "/policies"))
routes.append(("GET", "/policies"))
routes.append(("POST", "/processingunits"))
routes.append(("GET", "/processingunits"))
routes.append(("POST", "/servers"))
routes.append(("GET", "/servers"))
routes.append(("POST", "/systemcalls"))
routes.append(("GET", "/systemcalls"))
routes.append(("GET", "/tags"))
routes.append(("POST", "/users"))
routes.append(("GET", "/users"))

# routes for networkaccesspolicies
routes.append(("GET", "/networkaccesspolicies/:id"))
routes.append(("PUT", "/networkaccesspolicies/:id"))
routes.append(("DELETE", "/networkaccesspolicies/:id"))

