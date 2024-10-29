import { EventDescription } from './EventDescription';
import { EventOption } from './EventOption';
import { EventTarif } from './EventTarif';
import { Ticket } from './Ticket';
import { EventCategory } from './EventCategory';
import { EventTag } from './EventTag';

export interface Event {
  id: number;
  hostId: number;
  hostType: string;
  imageUrls: string[];
  videoUrl?: string;
  title: string;
  subtitle?: string;
  startDate: string;
  endDate: string;
  startTime: string;
  endTime: string;
  isOnline: boolean;
  isVisible: boolean;
  useStudibox: boolean;
  ticketPrice: number;
  ticketStock: number;
  address?: string;
  city?: string;
  postcode?: string;
  region?: string;
  country?: string;
  statistics?: string;
  ticketsSold: number;
  revenue: number;
  isValidatedByAdmin: boolean;
  descriptions: EventDescription[];
  options: EventOption[];
  tarifs: EventTarif[];
  tickets: Ticket[];
  categories: EventCategory[];
  tags: EventTag[];
  createdAt: string;
  updatedAt: string;
}
