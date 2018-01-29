from pyramid.view import view_config
from pyramid.httpexceptions import HTTPNotFound


@view_config(route_name='home', renderer='json')
def my_view(request):
    return {'project': 'lpmng-fw-agent'}
