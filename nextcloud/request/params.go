package request

const (
	listFileXML = `<?xml version="1.0"?>
	<d:propfind  xmlns:d="DAV:" xmlns:oc="http://owncloud.org/ns" xmlns:nc="http://nextcloud.org/ns">
	  <d:prop>
		   <d:getlastmodified />
		   <d:getetag />
		   <d:getcontenttype />
		   <d:resourcetype />
		   <oc:fileid />
		   <oc:permissions />
		   <oc:size />
		   <d:getcontentlength />
		   <nc:has-preview />
		   <oc:favorite />
		   <oc:comments-unread />
		   <oc:owner-display-name />
		   <oc:share-types />
	  </d:prop>
	</d:propfind>`
)
