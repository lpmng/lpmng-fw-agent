from pyramid.config import Configurator


def main(global_config, **settings):
    """ This function returns a Pyramid WSGI application.
    """
    config = Configurator(settings=settings)
    config.add_route('session_created_event', '/event/session/created')
    config.scan()
    return config.make_wsgi_app()
