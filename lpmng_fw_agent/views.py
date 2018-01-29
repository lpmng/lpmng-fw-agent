import logging

from pyramid.view import view_config
from pyramid.httpexceptions import HTTPNotFound
from pyramid.response import Response

log = logging.getLogger(__name__)


@view_config(route_name='session_created_event',
             renderer='json',
             request_method='POST')
def my_view(request):
    body = request.json_body
    data = body.get('param', {})

    mac = data.get('mac')
    ip4 = data.get('ip4')
    return Response(status=202)
